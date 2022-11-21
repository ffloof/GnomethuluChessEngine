package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

var DepthCount map[int8]int = map[int8]int{}

var TTable = make(TranspositionTable, 20*1024*1024, 20*1024*1024) //20MB * struct size
var Misses int

//Simplified minmax with ab pruning
func NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	DepthCount[depth] += 1

	entry := TTable.Get(board)
	if entry != nil {
		if depth <= entry.Depth {
			return entry.Score
		}
	}

	if depth == 0 {
		return CountMaterial(board)
	}

	moves := board.GenerateLegalMoves()
	if len(moves) == 0 {
		if board.OurKingInCheck() {
			return -9999
		}
		return 0
	}

	bestMove := moves[0]
	for _, move := range moves {
		undo := board.Apply(move)
		score := -NegaMax(board, -beta, -alpha, depth - 1)
		undo()

		if score > alpha {
			alpha = score
			if alpha >= beta {
				TTable.Set(board, bestMove, alpha, depth)
				return alpha
			}
		}
	} 
	TTable.Set(board, bestMove, alpha, depth)
	return alpha
}


func QuisenceSearch(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	
}