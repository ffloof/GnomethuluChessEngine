package mcts

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
)

func UCT(parent, child *MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 0.1
	return (child.Value / child.Visits) + math.Sqrt(c*math.Log(parent.Visits)/child.Visits)
}

// Development tables to encourage pieces to be on certain ranks

// Suggested on: https://www.chessprogramming.org/Simplified_Evaluation_Function
// With slight modifications
const pawnWeight float64 = 1.0
const knightWeight float64 = 3.2
const bishopWeight float64 = 3.3
const rookWeight float64 = 5.0
const queenWeight float64 = 9.0

//TODO: make boards symmetric so it doesnt inverse ranks
func reverse(s [64]float64) [64]float64{
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j]/100, s[i]/100
    }
    return s
}

// NOTE: indexing is in reverse so white pieces on top
var pawnDevelopment = reverse([64]float64{
	 0,  0,  0,  0,  0,  0,  0,  0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	 5,  5, 10, 25, 25, 10,  5,  5,
	 0,  0,  0, 20, 20,  0,  0,  0,
	 5, -5,-10,  0,  0,-10, -5,  5,
	 5, 10, 10,-20,-20, 10, 10,  5,
	 0,  0,  0,  0,  0,  0,  0,  0,
})
var knightDevelopment = reverse([64]float64{
	-50,-40,-30,-30,-30,-30,-40,-50,
	-40,-20,  0,  0,  0,  0,-20,-40,
	-30,  0, 10, 15, 15, 10,  0,-30,
	-30,  5, 15, 20, 20, 15,  5,-30,
	-30,  0, 15, 20, 20, 15,  0,-30,
	-30,  5, 10, 15, 15, 10,  5,-30,
	-40,-20,  0,  5,  5,  0,-20,-40,
	-50,-40,-30,-30,-30,-30,-40,-50,
})
var bishopDevelopment = reverse([64]float64{
	-20,-10,-10,-10,-10,-10,-10,-20,
	-10,  0,  0,  0,  0,  0,  0,-10,
	-10,  0,  5, 10, 10,  5,  0,-10,
	-10,  5,  5, 10, 10,  5,  5,-10,
	-10,  0, 10, 10, 10, 10,  0,-10,
	-10, 10, 10, 10, 10, 10, 10,-10,
	-10,  5,  0,  0,  0,  0,  5,-10,
	-20,-10,-10,-10,-10,-10,-10,-20,
})
var rookDevelopment = reverse([64]float64{
	  0,  0,  0,  0,  0,  0,  0,  0,
	  5, 10, 10, 10, 10, 10, 10,  5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	  0,  0,  0,  5,  5,  0,  0,  0,
})
var queenDevelopment = reverse([64]float64{
	-20,-10,-10, -5, -5,-10,-10,-20,
	-10,  0,  0,  0,  0,  0,  0,-10,
	-10,  0,  5,  5,  5,  5,  0,-10,
	 -5,  0,  5,  5,  5,  5,  0, -5,
	  0,  0,  5,  5,  5,  5,  0, -5,
	-10,  5,  5,  5,  5,  5,  0,-10,
	-10,  0,  5,  0,  0,  0,  0,-10,
	-20,-10,-10, -5, -5,-10,-10,-20,
})
var earlyKingDevelopment = reverse([64]float64{
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-20,-30,-30,-40,-40,-30,-30,-20,
	-10,-20,-20,-20,-20,-20,-20,-10,
	 20, 20,  0,  0,  0,  0, 20, 20,
	 20, 30, 10,  0,  0, 10, 30, 20,
})
var lateKingDevelopment = reverse([64]float64{
	-50,-40,-30,-20,-20,-30,-40,-50,
	-30,-20,-10,  0,  0,-10,-20,-30,
	-30,-10, 20, 30, 30, 20,-10,-30,
	-30,-10, 30, 40, 40, 30,-10,-30,
	-30,-10, 30, 40, 40, 30,-10,-30,
	-30,-10, 20, 30, 30, 20,-10,-30,
	-30,-30,  0,  0,  0,  0,-30,-30,
	-50,-30,-30,-30,-30,-30,-30,-50,
})

func Evaluate(board dragontoothmg.Board) float64 {
	eval := 0.0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				eval += pawnWeight
				eval += pawnDevelopment[i]
			}
			if board.White.Knights>>i%2 == 1 {
				eval += knightWeight
				eval += knightDevelopment[i]
			}
			if board.White.Bishops>>i%2 == 1 {
				eval += bishopWeight
				eval += bishopDevelopment[i]
			}
			if board.White.Rooks>>i%2 == 1 {
				eval += rookWeight
				eval += rookDevelopment[i]
			}
			if board.White.Queens>>i%2 == 1 {
				eval += queenWeight
				eval += queenDevelopment[i]

			}
			if board.White.Kings>>i%2 == 1 {
				eval += earlyKingDevelopment[i]
			}
		} else if board.Black.All>>i%2 == 1 {
			j := 63-i

			if board.Black.Pawns>>i%2 == 1 {
				eval -= pawnWeight
				eval -= pawnDevelopment[j]
			}
			if board.Black.Knights>>i%2 == 1 {
				eval -= knightWeight
				eval -= knightDevelopment[j]
			}
			if board.Black.Bishops>>i%2 == 1 {
				eval -= bishopWeight
				eval -= bishopDevelopment[j]
			}
			if board.Black.Rooks>>i%2 == 1 {
				eval -= rookWeight
				eval -= rookDevelopment[j]
			}
			if board.Black.Queens>>i%2 == 1 {
				eval -= queenWeight
				eval -= queenDevelopment[j]
			}
			if board.Black.Kings>>i%2 == 1 {
				eval -= earlyKingDevelopment[j]
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