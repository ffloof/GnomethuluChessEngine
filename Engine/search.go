package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type searcher struct {

}

func /*(search *searcher) */NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	//TODO: check if were in checkmate

	if depth == 0 {
		return MaterialCount(board)
	}

	moves := board.GenerateLegalMoves()
	if len(moves) == 0 {
		if board.OurKingInCheck() {
			return -10001
		} else {
			return 0
		}
	}

	for _, move := range moves {
		undo := board.Apply(move)
		score := -NegaMax(board,-beta,-alpha,depth -1)
		undo()

		if score > alpha {
			alpha = score
			if alpha >= beta {
				return score
			}
		}
	}
	return alpha
}