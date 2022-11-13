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
// TODO: try replacing with pesto eval
// With slight modifications
const pawnWeight float64 = 1.0
const knightWeight float64 = 3.0
const bishopWeight float64 = 3.2
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
var earlyPawnDevelopment = reverse([64]float64{
	 0,  0,  0,  0,  0,  0,  0,  0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 25, 25, 20, 10, 10,
	 5,  5, 10, 30, 30, 10,  5,  5,
	 0,  0,  0, 30, 30,  0,  0,  0,
	 5, -5,-10,  0,  0,-10, -5,  5,
	 5, 10, 10,-20,-20, 10, 10,  5,
	 0,  0,  0,  0,  0,  0,  0,  0,
})

var latePawnDevelopment = reverse([64]float64{
	 0,   0,  0,  0,  0,  0,  0,  0,
	 88, 99, 99, 99, 99, 99, 99, 88,
	 44, 49, 49, 49, 49, 49, 49, 44,
	 33, 28, 28, 28, 28, 28, 28, 33,
	  0,  5,  5,  5,  5,  5,  5,  0,
	-10, -5, -5, -5, -5, -5, -5,-10,
	-20,-15,-15,-15,-15,-15,-15,-20,
	  0,  0,  0,  0,  0,  0,  0,  0,
})



var knightDevelopment = reverse([64]float64{
	-50,-30,-20,-20,-20,-20,-30,-50,
	-40,-20,  0,  0,  0,  0,-20,-40,
	-30,  0, 10, 15, 15, 10,  0,-30,
	-30,  5, 15, 20, 20, 15,  5,-30,
	-30,  0, 15, 20, 20, 15,  0,-30,
	-30,  5, 10, 15, 15, 10,  5,-30,
	-40,-20,  0,  5,  5,  0,-20,-40,
	-50,-30,-20,-20,-20,-20,-30,-50,
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
	  5, 15, 15, 15, 15, 15, 15,  5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	 -5,  0,  0,  0,  0,  0,  0, -5,
	  0,  0,  5,  10, 10, 5, 0,  0,
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
	 20, 20,-10,-10,-10,-10, 20, 20,
	 20, 30, 10,-20,  0,-20, 30, 20,
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

	minorCount := 0
	queens := 0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				eval += pawnWeight
			} else if board.White.Knights>>i%2 == 1 {
				eval += knightWeight
				eval += knightDevelopment[i]
				minorCount += 1
			} else if board.White.Bishops>>i%2 == 1 {
				eval += bishopWeight
				eval += bishopDevelopment[i]
				minorCount += 1
			} else if board.White.Rooks>>i%2 == 1 {
				eval += rookWeight
				eval += rookDevelopment[i]
			} else if board.White.Queens>>i%2 == 1 {
				eval += queenWeight
				eval += queenDevelopment[i]
				queens += 1
			}
		} else if board.Black.All>>i%2 == 1 {
			j := 63-i

			if board.Black.Pawns>>i%2 == 1 {
				eval -= pawnWeight
			} else if board.Black.Knights>>i%2 == 1 {
				eval -= knightWeight
				eval -= knightDevelopment[j]
				minorCount += 1
			} else if board.Black.Bishops>>i%2 == 1 {
				eval -= bishopWeight
				eval -= bishopDevelopment[j]
				minorCount += 1
			} else if board.Black.Rooks>>i%2 == 1 {
				eval -= rookWeight
				eval -= rookDevelopment[j]
			} else if board.Black.Queens>>i%2 == 1 {
				eval -= queenWeight
				eval -= queenDevelopment[j]
				queens += 1
			}
			
		}
	}

	for i := 0; i < 64; i++ {
		j := 63-i
		if queens == 0 || minorCount < 2 {
			if board.White.Pawns>>i%2 == 1 {
				eval += latePawnDevelopment[i]
			} else if board.White.Kings>>i%2 == 1 {
				eval += lateKingDevelopment[i]
			} 

			if board.Black.Pawns>>i%2==1{
				eval -= latePawnDevelopment[j]
			} else if board.Black.Kings>>i%2 == 1 {
				eval -= lateKingDevelopment[j]
			}

		} else {
			if board.White.Pawns>>i%2 == 1 {
				eval += earlyPawnDevelopment[i]
			} else if board.White.Kings>>i%2 == 1 {
				eval += earlyKingDevelopment[i]
			} 

			if board.Black.Pawns>>i%2==1{
				eval -= earlyPawnDevelopment[j]
			} else if board.Black.Kings>>i%2 == 1 {
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
	a := 0.9
	return ((2*a) / (1 + math.Exp(-n*c))) - a
}