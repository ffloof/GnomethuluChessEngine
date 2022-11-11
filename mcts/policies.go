package mcts

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
	"fmt"
)

func UCT(parent, child *MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 0.3
	return (child.Value / child.Visits) + math.Sqrt(c*math.Log(parent.Visits)/child.Visits)
}

func Evaluate(board dragontoothmg.Board) float64 {
	eval := 0.0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				eval += 1.0
			}
			if board.White.Knights>>i%2 == 1 {
				eval += 3.0
			}
			if board.White.Bishops>>i%2 == 1 {
				eval += 3.0
			}
			if board.White.Rooks>>i%2 == 1 {
				eval += 5.0
			}
			if board.White.Queens>>i%2 == 1 {
				eval += 9.0
			}
		}
		if board.Black.All>>i%2 == 1 {
			if board.Black.Pawns>>i%2 == 1 {
				eval -= 1.0
				fmt.Println(i)
			}
			if board.Black.Knights>>i%2 == 1 {
				eval -= 3.0
			}
			if board.Black.Bishops>>i%2 == 1 {
				eval -= 3.0
			}
			if board.Black.Rooks>>i%2 == 1 {
				eval -= 5.0
			}
			if board.Black.Queens>>i%2 == 1 {
				eval -= 9.0
			}
		}
	}

	eval /= 15
	// TODO: make sure this doesnt lead to missing mate

	if board.Wtomove {
		eval = -eval
	}

	return eval
	//return SigmoidLike(eval)
}

// Very similar to a signmoid function except its on the range [-1,1]
func SigmoidLike(n float64) float64 {
	c := 1.0 // TODO: play around with tweaking c based off eval func
	return (2 / (1 + math.Exp(-n*c))) - 1
}