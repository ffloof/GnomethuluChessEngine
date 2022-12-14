package search

import (
	"github.com/ffloof/dragontoothmg"
)

func ControlMap(board *dragontoothmg.Board) *[64]int8 {
	threats := [64]int8{}
	ourMoves := board.GenerateControlMoves()
	board.Wtomove = !board.Wtomove
	opponentMoves := board.GenerateControlMoves()
	board.Wtomove = !board.Wtomove

	for _, move := range ourMoves {
		threats[move.To()] += 1
	}

	for _, move := range opponentMoves {
		threats[move.To()] -= 1
	}

	return &threats
}