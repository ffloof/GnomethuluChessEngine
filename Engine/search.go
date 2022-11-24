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
		Evaluation: Pesto,
	}
}

func max(a int16, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func min(a int16, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

//Simplified minmax with ab pruning
func (search *Searcher) NegaMax(board *dragontoothmg.Board, alpha, beta int16, depth int8) int16 {
	search.DepthCount[depth] += 1

	originalAlpha := alpha

	var tableMove dragontoothmg.Move = 0xFFFF
	entry := search.Table.Get(board, alpha, beta, depth)
	if entry != nil {
		if entry.Depth >= depth {
			if entry.Type == Exact { //Exact
				return entry.Score
			} else if entry.Type == Lowerbound { //Lowerbound
				if entry.Score > alpha {
					alpha = entry.Score
					if alpha >= beta {
						return entry.Score
					}
				}
			} else { //Upperbound
				if entry.Score < beta {
					beta = entry.Score
					if alpha >= beta {
						return entry.Score
					}
				}
			}
		} 
		tableMove = entry.BestMove
	}

	order := CreateMoveOrder(board, tableMove, depth <= 0)

	if order.NoMoves() {
		if board.OurKingInCheck() {
			return -10000
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



	

	var bestScore int16 = -10000
	var bestMove dragontoothmg.Move = 0xFFFF

	for {
		move, done := order.GetNextMove()
		if done { break }

		undo := board.Apply(move)
		score := -search.NegaMax(board, -beta, -alpha, depth - 1)
		undo()

		//if board.ToFen() == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" {
			//fmt.Println(move.String(), score)
		//}


		if score > bestScore {
			bestScore = score
			bestMove = move

			if score > alpha {
				alpha = score
				if alpha >= beta {
					break
				}
			}
		}
	}

	search.Table.Set(board, bestMove, bestScore, originalAlpha, beta, depth)
	return bestScore
}