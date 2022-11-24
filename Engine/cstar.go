package engine

import (
	"fmt"
	"time"
	"github.com/dylhunn/dragontoothmg"
)

var MINTIME int64 = 3000

/* CSTAR cant really work efficiently or accurately unless we fix the transposition table for null windows
func Cstar(board dragontoothmg.Board){
	searcher := NewSearch()
	start := time.Now()

	var roughScore int16
	var depth int8 = 2

	for depth < 100 {
		var lowerbound int16 = -9999
		var upperbound int16 = 9999
		var roughness int16 = 2

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
			roughScore = (lowerbound + upperbound) / 2
			searcher.NegaMax(&board, lowerbound-1, lowerbound, depth)
			break
		}
		depth++
	}

	fmt.Println("CSTAR", searcher.Table.Get(&board).BestMove.String(), roughScore, depth)
}*/

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
	fmt.Println("BASE",searcher.Table.Get(&board).BestMove.String(), score, depth)
	searcher.PrintDepths(depth)
}