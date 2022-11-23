package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

//TODO: create another implementation where it just simply sorts moves, benchmark accordingly
type moveOrder struct {
	moves []dragontoothmg.Move
	scores map[dragontoothmg.Move]int
	startIndex int
}

func CreateMoveOrder(board *dragontoothmg.Board, tableMove dragontoothmg.Move, capturesOnly bool) *moveOrder {
	moves := board.GenerateLegalMoves()
	scores := map[dragontoothmg.Move]int{}
	
	for _, move := range moves {
		scores[move] = scoreMove(board, move, tableMove, capturesOnly)
	}

	return &moveOrder {
		moves: moves,
		scores: scores,
		startIndex: 0,
	}
}

func (order *moveOrder) GetNextMove() (dragontoothmg.Move,bool) {
	bestIndex := -1
	bestScore := -1

	for i := order.startIndex; i < len(order.moves); i++ {
		if order.scores[order.moves[i]] > bestScore {
			bestIndex = i
		}
	}

	if bestIndex < 0 {
		order.startIndex = len(order.moves)
		return 0, true
	}
	
	chosenMove := order.moves[bestIndex]
	order.moves[order.startIndex], order.moves[bestIndex] = order.moves[bestIndex], order.moves[order.startIndex]
	order.startIndex++
	return chosenMove, false
}



func scoreMove (board *dragontoothmg.Board, move dragontoothmg.Move, tableMove dragontoothmg.Move, capturesOnly bool) int {
	if dragontoothmg.IsCapture(move, board) {
		return 0
	}

	if capturesOnly {
		return -1
	} else {
		return 1
	}
}

func (order *moveOrder) NoMoves() bool {
	return len(order.moves) == 0
}