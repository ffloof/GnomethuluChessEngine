package main

import (
	"fmt"
	"gnomethulu/evaluation/traditional"
	"gnomethulu/policy"
	"gnomethulu/uci"
	"gnomethulu/search"
	"github.com/dylhunn/dragontoothmg"
	//"gnomethulu/evaluation/neural"
)

func main() {	
	tournament()
	uci.Init(policy.MVVLVA_UCT, traditional.CustomV1)
}

// Runs a torunament between two ai configurations
func tournament(){
	searcher1 := search.NewSearch(policy.MVVLVA_UCT, traditional.CustomV1)
	searcher2 := search.NewSearch(policy.UCT, traditional.CustomV1)
	searcher1.PolicyExplore = 0.15
	searcher1.EmptyVisits = 0.5
	searcher2.PolicyExplore = 0.1
	searcher2.EmptyVisits = 0.01

	handicap := 0.5

	startFens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", //No Moves
		"r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 3 3", //Ruey Lopez
		"rnbqkbnr/ppp1pppp/8/8/2pPP3/8/PP3PPP/RNBQKBNR b KQkq - 0 3", //Queens Gambit Accepted Full Center
		"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/8/PP2PPPP/RNBQKBNR w KQkq - 0 3", //Queens Gambit Declined
		"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/8/PP2PPPP/RNBQKBNR w KQkq - 0 3", //Slav Defense
		"r1bqkbnr/pp1ppppp/2n5/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3", //Open Sicilian Nc6
		"rnbqkbnr/pp2pppp/3p4/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 3", //Open Sicilian d6
		"rnbqkbnr/pp2pppp/2p5/3p4/3PP3/8/PPP2PPP/RNBQKBNR w KQkq - 0 3", //Caro Kann
		"rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/8/PPP2PPP/RNBQKBNR w KQkq - 0 3", //French Defense
		"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N5/PP2PPPP/R1BQKBNR w KQkq - 3 4", //Nimzo Indian
		"rnbqkb1r/ppp1pppp/3p1n2/8/3PP3/8/PPP2PPP/RNBQKBNR w KQkq - 1 3", //Pirc Defense
		"rnbqk2r/pppp1ppp/5n2/4B3/1b6/8/P1PPPPPP/RN1QKBNR w KQkq - 1 4", //Orangutan Opening Exchange Variation
		"rnbqkbnr/pppp1ppp/8/8/2B1P3/8/PB3PPP/RN1QK1NR b KQkq - 0 5", //Danish Gambit
		"rnbqkbnr/pppp1ppp/8/4p3/4PP2/8/PPPP2PP/RNBQKBNR b KQkq - 0 2", //Kings Gambit
		"rnbqkb1r/pppp1ppp/5n2/4p3/4PP2/2N5/PPPP2PP/R1BQKBNR b KQkq - 0 3", //Vienna Gambit
		"rnbqkbnr/pp1ppppp/8/2p5/3P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 2", //Old Benoni
		"rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2", //Scandinavian Defense
		"rnbqkbnr/ppppp1pp/8/5p2/3P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 2", //Dutch Defense
		"rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1", //Anderssen's Opening
		"rnbqkbnr/pppppppp/8/8/P7/8/1PPPPPPP/RNBQKBNR b KQkq - 0 1", //Ware Opening
		"rnbqkb1r/pppppppp/5n2/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 1 2", //Alekhine's Defense
		"rnbqkbnr/pp1ppppp/8/2p5/3PP3/8/PPP2PPP/RNBQKBNR b KQkq - 0 2", //Smith-Morra Gambit
		"rnbqkbnr/ppppp1pp/8/5p2/3PP3/8/PPP2PPP/RNBQKBNR b KQkq - 0 2", //Staunton Gambit
		"rnbqkbnr/pppp1ppp/8/4p3/2P5/2N5/PP1PPPPP/R1BQKBNR b KQkq - 1 2", //King's English
		"rnbqkbnr/ppp1pppp/8/3p4/3P1B2/8/PPP1PPPP/RN1QKBNR b KQkq - 1 2", //London System
		"rnbqkb1r/pppp1ppp/5n2/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3", //Petrov Defense
		"r1bqkbnr/pppp1ppp/2n5/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq - 0 3", //Scotch Game
		"rnbqkb1r/pppp1ppp/8/4p3/2B1n3/2N5/PPPP1PPP/R1BQK1NR w KQkq - 0 4", //Frankenstein-Dracula Variation
		"r1bqk2r/pppp1Bpp/2n2n2/2b1p1N1/4P3/8/PPPP1PPP/RNBQK2R b KQkq - 0 5", //Traxler
		"rnbqkbnr/pppppppp/8/8/7P/8/PPPPPPP1/RNBQKBNR b KQkq - 0 1", //Desprez/Kadas Opening
		//TODO: add englund and black mardiemer?
	}

	advantage := 0.0
	for i, opening := range startFens {
		advantage += runGame(searcher1, searcher2, opening, handicap)
		advantage -= runGame(searcher2, searcher1, opening, 1/handicap)
		fmt.Println("Games:", 2 * (i+1), "Advantage:", advantage)
	}
}

// Returns 1 if white wins, 0.5 if draw, 0 if black wins
func runGame(whiteSearcher, blackSearcher *search.MonteCarloTreeSearcher, fen string, handicap float64) float64 {
	board := dragontoothmg.ParseFen(fen)
	moveCounter := 0
	for {
		if moveCounter > 150 {
			return 0.5
		}
		moveCounter++

		if len(board.GenerateLegalMoves()) == 0 {
			if board.OurKingInCheck() {
				if board.Wtomove {
					return 0
				} else {
					return 1
				}
			}
			return 0.5
		}

		whiteSearcher.SetPosition(board)
		blackSearcher.SetPosition(board)

		var bestMove dragontoothmg.Move
		if board.Wtomove {
			whiteSearcher.RunIterations(int(10000*handicap))
			bestMove = whiteSearcher.GetBestMove()
		} else {
			blackSearcher.RunIterations(10000)
			bestMove = blackSearcher.GetBestMove()
		}
		board.Apply(bestMove)
	}
}