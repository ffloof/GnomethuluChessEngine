package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Searcher struct {
	Table TranspositionTable
	DepthCount map[int8]int
	Evaluation func(board *dragontoothmg.Board)int16
}

func NewSearch() *Searcher {
	return &Searcher {
		Table: make(TranspositionTable, 20*1024*1024), // 20MB * struct size
		DepthCount: map[int8]int{},
		Evaluation: CountMaterial,
	}
}

//Simplified minmax with ab pruning
func (search *Searcher) NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	search.DepthCount[depth] += 1

	var tableMove dragontoothmg.Move
	entry := search.Table.Get(board)
	if entry != nil {
		if depth <= entry.Depth {
			return entry.Score
		} else {
			tableMove = entry.BestMove
		}
	}

	order := CreateMoveOrder(board, tableMove, depth <= 0)

	if order.NoMoves() {
		if board.OurKingInCheck() {
			return -9999
		}
		return 0
	}

	if depth <= 0 {
		score := search.Evaluation(board)

		if score >= alpha {
			alpha = score
			if alpha >= beta {
				return alpha
			}
		}
	}

	var bestMove dragontoothmg.Move
	for {
		move, done := order.GetNextMove()
		if done { break }

		undo := board.Apply(move)
		score := -search.NegaMax(board, -beta, -alpha, depth - 1)
		undo()

		if score > alpha {
			alpha = score
			bestMove = move
			if alpha >= beta {
				break
			}
		}
	}

	search.Table.Set(board, bestMove, alpha, depth)
	return alpha
}