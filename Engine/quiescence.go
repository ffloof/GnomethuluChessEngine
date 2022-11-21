package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

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
		if !dragontoothmg.IsCapture(move, board) { continue }

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

// Includes everything until either player is happy
func FullQuiescenceSearch(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
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