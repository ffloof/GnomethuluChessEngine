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
		return QuisenceSearch(board, alpha, beta, depth)
		//return CountMaterial(board)
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


func QuiescenceSearch(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	DepthCount[depth] += 1

	entry := TTable.Get(board)
	if entry != nil {
		if depth <= entry.Depth {
			return entry.Score
		}
	}

	moves := board.GenerateLegalMoves()
	
	if len(moves) == 0 {
		if board.OurKingInCheck() {
			return -1.0
		} else {
			return 0.0
		}
	}

	score := CountMaterial(board)

	if score >= alpha {
		alpha = score
		if alpha >= beta {
			return alpha
		}
	}

	bestMove := moves[0]
	for _, move := range moves {
		undo := board.Apply(move) 
		score := -QuiescenceSearch(board, -beta, -alpha, depth - 1)
		undo()

		if score >= alpha {
            alpha = score   
            bestMove = move
            if alpha >= beta {
            	break
			}  
        }
	}
	//TODO: consider removing tbh
	TTable.Set(board, bestMove, alpha, depth)
	return alpha
}