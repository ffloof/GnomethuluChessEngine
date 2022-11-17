package evaluation

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
	"sort"
)

// Tables for pesto eval
//TODO: see if its worth creating two arrays one inversed and one not so that it doesnt have to do i^56 at runtime
func reverse(fixed int, s [64]int) [64]int{
	var newS [64]int
	for i := 0; i < 64; i++ {
		j := i ^ 56
		newS[j] = fixed+s[i]
	}
	return newS
}

var mgPawnTable [64]int = reverse(82, [64]int{
	  0,   0,   0,   0,   0,   0,  0,   0,
	 98, 134,  61,  95,  68, 126, 34, -11,
	 -6,   7,  26,  31,  65,  56, 25, -20,
	-14,  13,  10,  21,  23,  12, 17, -23,
	-30,  -2,   5,  12,  17,   6, -5, -30,
	-29,  -4,  -4, -14, -10,   3, 33, -12,
	-35,  -1, -15, -25, -30,  24, 38, -10,
	  0,   0,   0,   0,   0,   0,  0,   0,
})

var egPawnTable [64]int = reverse(94,[64]int{
	  0,   0,   0,   0,   0,   0,   0,   0,
	178, 173, 158, 134, 147, 132, 165, 187,
	 94, 100,  85,  67,  56,  53,  82,  84,
	 32,  24,  13,   5,  -2,   4,  17,  17,
	 13,   9,  -3,  -7,  -7,  -8,   3,  -1,
	  4,   7,  -6,   1,   0,  -5,  -1,  -8,
	 13,   8,   8,  10,  13,   0,   2,  -7,
	  0,   0,   0,   0,   0,   0,   0,   0,
})

var mgKnightTable [64]int = reverse(337, [64]int{
	-167, -89, -34, -49,  61, -97, -15, -107,
	 -73, -41,  72,  36,  23,  62,   7,  -17,
	 -47,  60,  37,  65,  84, 129,  73,   44,
	  -9,  17,  19,  20,  15,  69,  18,   22,
	 -13,   4,  16,  13,  28,  19,  21,   -8,
	 -23,  -9,  12,  10,  19,  17,  25,  -16,
	 -29, -53, -12,  -3,  -1,  18, -14,  -19,
	-105, -21, -58, -33, -17, -28, -19,  -23,
})

var egKnightTable [64]int = reverse(281,[64]int{
	-58, -38, -13, -28, -31, -27, -63, -99,
	-25,  -8, -25,  -2,  -9, -25, -24, -52,
	-24, -20,  10,   9,  -1,  -9, -19, -41,
	-17,   3,  22,  22,  22,  11,   8, -18,
	-18,  -6,  16,  25,  16,  17,   4, -18,
	-23,  -3,  -1,  15,  10,  -3, -20, -22,
	-42, -20, -10,  -5,  -2, -20, -23, -44,
	-29, -51, -23, -15, -22, -18, -50, -64,
})

var mgBishopTable [64]int = reverse(365, [64]int{
	-29,   4, -82, -37, -25, -42,   7,  -8,
	-26,  16, -18, -13,  30,  59,  18, -47,
	-16,  37,  43,  40,  35,  50,  37,  -2,
	 -4,   9,  19,  20,  22,  37,   5,  -2,
	 -6,  13,  18,  30,  34,  20,  10,  -4,
	  0,  12,  15,  10,  -5,  27,  12,  10,
	  4,  15,  16,   5,   7,  21,  33,   1,
	-33,  -3, -14, -21, -13, -15, -39, -21,
})

var egBishopTable [64]int = reverse(297,[64]int{
	-14, -21, -11,  -8, -7,  -9, -17, -24,
	 -8,  -4,   7, -12, -3, -13,  -4, -14,
	  2,  -8,   0,  -1, -2,   6,   0,   4,
	 -3,   9,  12,   9, 14,  10,   3,   2,
	 -6,   3,  13,  19,  7,  10,  -3,  -9,
	-12,  -3,   8,  10, 13,   3,  -7, -15,
	-14, -18,  -7,  -1,  4,  -9, -15, -27,
	-23,  -9, -23,  -5, -9, -16,  -5, -17,
})

var mgRookTable [64]int = reverse(447,[64]int{
	 32,  42,  32,  51, 63,  9,  31,  43,
	 27,  32,  58,  62, 80, 67,  26,  44,
	 -5,  19,  26,  36, 17, 45,  61,  16,
	-24, -11,   7,  26, 24, 35,  -8, -20,
	-36, -26, -12,  -1,  9, -7,   6, -23,
	-45, -25, -16, -17,  3,  0,  -5, -33,
	-44, -16, -20,  -9, -1, 11,  -6, -71,
	-19, -13,  -5,  15, 12,  7, -37, -26,
})

var egRookTable [64]int = reverse(512,[64]int{
	13, 10, 18, 15, 12,  12,   8,   5,
	11, 13, 13, 11, -3,   3,   8,   3,
	 7,  7,  7,  5,  4,  -3,  -5,  -3,
	 4,  3, 13,  1,  2,   1,  -1,   2,
	 3,  5,  8,  4, -5,  -6,  -8, -11,
	-4,  0, -5, -1, -7, -12,  -8, -16,
	-6, -6,  0,  2, -9,  -9, -11,  -3,
	-9,  2,  3, -1, -5, -13,   4, -20,
})


var mgQueenTable [64]int = reverse(1025,[64]int{
	-28,   0,  29,  12,  59,  44,  43,  45,
	-24, -39,  -5,   1, -16,  57,  28,  54,
	-13, -17,   7,   8,  29,  56,  47,  57,
	-27, -27, -16, -16,  -1,  17,  -2,   1,
	 -9, -26,  -9, -10,  -2,  -4,   3,  -3,
	-14,   2, -11,  -2,  -5,   2,  14,   5,
	-35,  -8,  11,   2,   8,  15,  -3,   1,
	 -1, -18,  -9,  -5, -15, -25, -31, -50,
})

var egQueenTable [64]int = reverse(936,[64]int{
	 -9,  22,  22,  27,  27,  19,  10,  20,
	-17,  20,  32,  41,  58,  25,  30,   0,
	-20,   6,   9,  49,  47,  35,  19,   9,
	  3,  22,  24,  45,  57,  40,  57,  36,
	-18,  28,  19,  47,  31,  34,  39,  23,
	-16, -27,  15,   6,   9,  17,  10,   5,
	-22, -23, -30, -16, -16, -23, -36, -32,
	-33, -28, -22, -43,  -5, -32, -20, -41,
})

var mgKingTable [64]int = reverse(0,[64]int{
	-65,  23,  16, -15, -56, -34,   2,  13,
	 29,  -1, -20,  -7,  -8,  -4, -38, -29,
	 -9,  24,   2, -16, -20,   6,  22, -22,
	-17, -20, -12, -27, -30, -25, -14, -36,
	-49,  -1, -27, -39, -46, -44, -33, -51,
	-14, -14, -22, -46, -44, -30, -15, -27,
	  1,   7,  -8, -64, -43, -16,   9,   8,
	-15,  36,  12, -54,   8, -28,  24,  14,
})

var egKingTable [64]int = reverse(0,[64]int{
	-74, -35, -18, -18, -11,  15,   4, -17,
	-12,  17,  14,  17,  17,  38,  23,  11,
	 10,  17,  23,  15,  20,  45,  44,  13,
	 -8,  22,  24,  27,  26,  33,  26,   3,
	-18,  -4,  21,  24,  27,  23,   9, -11,
	-19,  -3,  11,  21,  23,  16,   7,  -9,
	-27, -11,   4,  13,  14,   4,  -5, -17,
	-53, -34, -21, -11, -28, -14, -24, -43,
})


func Wrapper(board dragontoothmg.Board) float64 {
	return PestoQuiescence(board, -0.9, 0.9)
}

// Consider transposition table and adding checks
func PestoQuiescence(board dragontoothmg.Board, alpha, beta float64) float64 {
	all_moves := board.GenerateLegalMoves()
	
	if len(all_moves) == 0 {
		if board.OurKingInCheck() {
			return -1.0
		} else {
			return 0.0
		}
	}

	score := Pesto(board)
	if score >= beta {
		return score
	}

	if score >= alpha {
		alpha = score
	}
	
	var chosen_moves []dragontoothmg.Move
	promote_moves := []dragontoothmg.Move{} 
	capture_moves := []dragontoothmg.Move{}

	if board.OurKingInCheck() {
		chosen_moves = all_moves
	} else {
		for _, move := range all_moves {
			promotePiece := move.Promote()
			if promotePiece == dragontoothmg.Nothing {
				if dragontoothmg.IsCapture(move, &board) {
					capture_moves = append(chosen_moves, move)
				} 
			} else if promotePiece == dragontoothmg.Queen { //Queen
				promote_moves = append(promote_moves, move)
			}
		}

		
		Less_MVV_LVA := func(c, d int) bool{
			a := capture_moves[c]
			b := capture_moves[d]

			victimAType, _ := dragontoothmg.GetPieceType(a.To(), &board)
			victimBType, _ := dragontoothmg.GetPieceType(b.To(), &board)

			if victimAType != victimBType  {
				return victimAType > victimBType
			} else {
				attackerAType, _ := dragontoothmg.GetPieceType(a.From(), &board)
				attackerBType, _ := dragontoothmg.GetPieceType(b.From(), &board)
				return attackerAType < attackerBType
			}
		}

		sort.Slice(chosen_moves, Less_MVV_LVA)

		chosen_moves = append(promote_moves, capture_moves...)
	}

	for _, move := range chosen_moves {
		if dragontoothmg.IsCapture(move, &board) {
			undo := board.Apply(move) 
			
			score = -PestoQuiescence(board, -beta, -alpha)
			
			undo()

			if score >= alpha {
                alpha = score   
                if alpha >= beta {
                	break
				}  
            }
		}
	}
	return alpha
}


//TODO: convert stuff in eval to int operations, with cast to float at end and compare speed
func Pesto(board dragontoothmg.Board) float64 {
	phase := 0
	midScore := 0
	endScore := 0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				midScore += mgPawnTable[i]
				endScore += egPawnTable[i]
			} else if board.White.Knights>>i%2 == 1 {
				midScore += mgKnightTable[i]
				endScore += egKnightTable[i]
				phase += 1
			} else if board.White.Bishops>>i%2 == 1 {
				midScore += mgBishopTable[i]
				endScore += egBishopTable[i]
				phase += 1
			} else if board.White.Rooks>>i%2 == 1 {
				midScore += mgRookTable[i]
				endScore += egRookTable[i]
				phase += 2
			} else if board.White.Queens>>i%2 == 1 {
				midScore += mgQueenTable[i]
				endScore += egQueenTable[i]
				phase += 4
			} else if board.White.Kings>>i%2 == 1{
				midScore += mgKingTable[i]
				endScore += egKingTable[i]
			}
		} else if board.Black.All>>i%2 == 1 {
			j := i ^ 56
			if board.Black.Pawns>>i%2 == 1 {
				midScore -= mgPawnTable[j]
				endScore -= egPawnTable[j]
			} else if board.Black.Knights>>i%2 == 1 {
				midScore -= mgKnightTable[j]
				endScore -= egKnightTable[j]
				phase += 1
			} else if board.Black.Bishops>>i%2 == 1 {
				midScore -= mgBishopTable[j]
				endScore -= egBishopTable[j]
				phase += 1
			} else if board.Black.Rooks>>i%2 == 1 {
				midScore -= mgRookTable[j]
				endScore -= egRookTable[j]
				phase += 2
			} else if board.Black.Queens>>i%2 == 1 {
				midScore -= mgQueenTable[j]
				endScore -= egQueenTable[j]
				phase += 4
			} else if board.Black.Kings>>i%2 == 1{
				midScore -= mgKingTable[j]
				endScore -= egKingTable[j]
			}
		}
	}


	if phase > 24 {
		phase = 24
	}

	eval := float64((midScore * phase) + (endScore * (24-phase)))/24/100 

	if !board.Wtomove {
		eval = -eval
	}

	return SigmoidLike(eval)
}

var SigmoidCurve float64 = 0.5
var SigmoidScale float64 = 0.9


// Very similar to a signmoid function except its on the range [-scale,scale]
// Play with c and a parameters
func SigmoidLike(n float64) float64 {
	n *= SigmoidCurve
	return ((2*SigmoidScale) / (1 + math.Exp(-n))) - SigmoidScale
}

