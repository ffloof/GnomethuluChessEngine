package engine

import (
	"fmt"
	"time"
	"github.com/dylhunn/dragontoothmg"
)

var MINTIME int64 = 3000

// CSTAR cant really work efficiently or accurately unless we fix the transposition table for null windows
func Cstar(board dragontoothmg.Board){
	searcher := NewSearch()
	start := time.Now()

	var depth int8 = 2

	for depth < 100 {
		var lowerbound int16 = -9999
		var upperbound int16 = 9999
		var roughness int16 = 2
		fmt.Println(depth)

		for lowerbound <= upperbound - roughness {

			gamma := (lowerbound + upperbound + 1)/2
			score := searcher.NegaMax(&board, gamma-1, gamma, depth)

			if score >= gamma {
				lowerbound = score
			}
			if score < gamma {
				upperbound = score
			}
		}
		
		if time.Since(start).Milliseconds() > MINTIME {
			searcher.NegaMax(&board, lowerbound-1, lowerbound, depth)
			entry := searcher.Table.Get(&board,lowerbound-1,lowerbound,depth)
			fmt.Println("CSTAR", entry.BestMove.String(), entry.Score, depth)
			break
		}
		depth++
	}

}

func Base(board dragontoothmg.Board){
	searcher := NewSearch()
	start := time.Now()
	
	var depth int8 = 2
	var score int16
	for depth < 100 {
		fmt.Println(depth)
		score = searcher.NegaMax(&board, -9999, 9999, depth)

		if time.Since(start).Milliseconds() > MINTIME {
			break
		}
		depth++
	}

	info := searcher.Table.Get(&board,-9999,9999,depth)
	fmt.Println("BASE",info.BestMove.String(), score, depth)
	searcher.PrintDepths(depth)
}

func VerifyBase(board dragontoothmg.Board){
	searcher := NewSearch()
	
	var depth int8 = 8
	moves := board.GenerateLegalMoves()

	for _, move := range moves {
		undo := board.Apply(move) 
		score := -searcher.NegaMax(&board, -9999, 9999, depth-1)
		undo()
		fmt.Println(move.String(), score)
	}
}