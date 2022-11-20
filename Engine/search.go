package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Searcher struct {
	Table TranspositionTable
}

func (search *Searcher) NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	entry := search.Table.Get(board)
	if entry != nil {
		if depth <= entry.Depth {
			return entry.Score
		} else {
			//TODO: set best move to be tried first somehow
		}
	}

	if depth == 0 {
		//TODO: do quiesence search
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

	var bestmove dragontoothmg.Move
	for _, move := range moves {
		undo := board.Apply(move)
		score := -search.NegaMax(board,-beta,-alpha,depth -1)
		undo()

		if score > alpha {
			alpha = score
			bestmove = move
			if alpha >= beta {
				search.Table.Set(board,depth,score,bestmove)
				return score
			}
		}
	}
	search.Table.Set(board,depth,alpha,bestmove)
	return alpha
}