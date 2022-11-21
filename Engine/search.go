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
		return QuiescenceSearch(board, alpha, beta, depth)
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
			bestMove = move
			if alpha >= beta {
				break
			}
		}
	} 
	TTable.Set(board, bestMove, alpha, depth)
	return alpha
}

//TODO: benchmark variations that only search captures, and use checks to make sure invalid scores arent used
