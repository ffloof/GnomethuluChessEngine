package search

import (
	"fmt"
	"strconv"
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

func PrintControlMap(position *[64]int8){
	for y := 7; y >= 0; y-- {
		line := ""
		for x := 0; x < 8; x++ {
			if position[y*8+x] >= 0 {
				line += " "
			}
			line += strconv.Itoa(int(position[y*8+x])) + " "
		}
		fmt.Println(line)
	}
}