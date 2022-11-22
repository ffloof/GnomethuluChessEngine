package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Searcher struct {
	Table TranspositionTable
	DepthCount map[int8]int
}

func NewSearch() *Searcher {
	return &Searcher {
		Table: make(TranspositionTable, 20*1024*1024), // 20MB * struct size
		DepthCount: map[int8]int{},
	}
}

//Simplified minmax with ab pruning
func (search *Searcher) NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	search.DepthCount[depth] += 1

	entry := search.Table.Get(board)
	if entry != nil {
		if depth <= entry.Depth {
			return entry.Score
		}
	}

	if depth == 0 {
		return search.QuiescenceSearch(board, alpha, beta, depth-1)
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

//TODO: benchmark variations that only search captures, and use checks to make sure invalid scores arent used
