package tournament

import (
	"fmt"
	"gnomethulu/search"
	"github.com/ffloof/dragontoothmg"
	"gnomethulu/evaluation/traditional"
	"gnomethulu/policy"
)

// INVESTIGATE: We could make the ai play the game from both sides until the ais disagree a move
// From here we can pit the ais, each playing their prefered move and recording which scored better, as well as maybe also some stockfish data



// This tournament pits two versions of the ai agains each other
// It plays from a variety of opening positions with each side playing once as white and once as black
// It records the score to see which ai is stronger

var openings = map[string]string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": "Starting Board",
	"rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2": "Scandinavian Defense",
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2": "Open Sicilian",
	"r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 1 3": "Ruy Lopez",
	"rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b KQkq - 0 2": "Queens Gambit",
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N5/PP2PPPP/R1BQKBNR w KQkq - 3 4": "Nimzo Indian",
}


func Run(){
	runTournament(
		search.NewSearch(policy.HeurUCT, traditional.CustomV2),
		search.NewSearch(policy.UCT, traditional.CustomV2),
	)
}

func runTournament(searcher1, searcher2 *search.MonteCarloTreeSearcher){
	scoreline := 0.0
	for fen, name := range openings {
		scoreline += playRound(searcher1, searcher2, fen)
		scoreline -= playRound(searcher2, searcher1, fen)
		fmt.Println(name, scoreline)
	}
}

func playRound(whitePlayer, blackPlayer *search.MonteCarloTreeSearcher, startFen string) float64 {
	position := dragontoothmg.ParseFen(startFen)
	maxmoves := 200
	for {
		maxmoves--
		if maxmoves < 0 {
			return 0.5 //Draw if too many moves
		}

		var nextMove dragontoothmg.Move
		if position.Wtomove {
			whitePlayer.SetPosition(position)
			whitePlayer.RunIterations(10000)
			nextMove = whitePlayer.GetBestMove()
		} else {
			blackPlayer.SetPosition(position)
			blackPlayer.RunIterations(10000)
			nextMove = blackPlayer.GetBestMove()
		}
		position.Apply(nextMove)
		moves := position.GenerateLegalMoves()
		if len(moves) == 0 {
			if position.OurKingInCheck() {
				if position.Wtomove {
					return 0 // White lost
				} else {
					return 1 // Black won
				}
			} else {
				return 0.5 // Stalemate
			}
		} else if position.White.All | position.Black.All == position.White.Kings | position.Black.Kings {
			// Checks if only the kings are on the board, ie draw by insufficient material
			return 0.5
		}
	}
}