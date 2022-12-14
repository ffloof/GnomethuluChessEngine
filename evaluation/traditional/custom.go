package traditional

import (
	"github.com/ffloof/dragontoothmg"
)

//TODO: Using bitboards its trivial to detect passed pawns
// Something similar to passed pawns can be used to detect outposts

// Tables for custom eval
var earlyPawnTable [64]int = reverse(82, [64]int{
	  0,   0,   0,   0,   0,   0,  0,   0,
	 98, 134,  61,  95,  68, 126, 34, -11,
	 -6,   7,  26,  31,  65,  56, 25, -20,
	-14,  13,  10,  21,  23,  12, 17, -23,
	-30,  -2,   5,  12,  17,   6, -5, -30,
	-29,  -4,  -4,  -7,  -3,   3, 33, -12,
	-35,  -1, -15, -25, -30,  24, 38, -10,
	  0,   0,   0,   0,   0,   0,  0,   0,
})

var latePawnTable [64]int = reverse(94,[64]int{
	  0,   0,   0,   0,   0,   0,   0,   0,
	178, 173, 158, 134, 147, 132, 165, 187,
	 94, 100,  85,  67,  56,  53,  82,  84,
	 32,  24,  13,   5,  -2,   4,  17,  17,
	 13,   9,  -3,  -7,  -7,  -8,   3,  -1,
	  4,   7,  -6,   1,   0,  -5,  -1,  -8,
	 13,   8,   8,  10,  13,   0,   2,  -7,
	  0,   0,   0,   0,   0,   0,   0,   0,
})

var earlyKnightTable [64]int = reverse(337, [64]int{
	-167, -89, -34, -49,  61, -97, -15, -107,
	 -73, -41,  72,  36,  23,  62,   7,  -17,
	 -47,  60,  37,  65,  84, 129,  73,   44,
	  -9,  17,  19,  20,  15,  69,  18,   22,
	 -13,   4,  16,  13,  28,  19,  21,   -8,
	 -23,  -9,  12,  10,  19,  17,  25,  -16,
	 -29, -53, -12,  -3,  -1,  18, -14,  -19,
	-105, -21, -58, -33, -17, -28, -19,  -23,
})

var lateKnightTable [64]int = reverse(281,[64]int{
	-58, -38, -13, -28, -31, -27, -63, -99,
	-25,  -8, -25,  -2,  -9, -25, -24, -52,
	-24, -20,  10,   9,  -1,  -9, -19, -41,
	-17,   3,  22,  22,  22,  11,   8, -18,
	-18,  -6,  16,  25,  16,  17,   4, -18,
	-23,  -3,  -1,  15,  10,  -3, -20, -22,
	-42, -20, -10,  -5,  -2, -20, -23, -44,
	-29, -51, -23, -15, -22, -18, -50, -64,
})

var earlyBishopTable [64]int = reverse(365, [64]int{
	-29,   4, -82, -37, -25, -42,   7,  -8,
	-26,  16, -18, -13,  30,  59,  18, -47,
	-16,  37,  43,  40,  35,  50,  37,  -2,
	 -4,   9,  19,  20,  22,  37,   5,  -2,
	 -6,  13,  18,  30,  34,  20,  10,  -4,
	  0,  12,  15,  -5, -10,  27,  12,  10,
	  4,  15,  16,   8,  12,  21,  33,   1,
	-33,  -3, -19, -21, -13, -20, -39, -21,
})

var lateBishopTable [64]int = reverse(297,[64]int{
	-14, -21, -11,  -8, -7,  -9, -17, -24,
	 -8,  -4,   7, -12, -3, -13,  -4, -14,
	  2,  -8,   0,  -1, -2,   6,   0,   4,
	 -3,   9,  12,   9, 14,  10,   3,   2,
	 -6,   3,  13,  19,  7,  10,  -3,  -9,
	-12,  -3,   8,  10, 13,   3,  -7, -15,
	-14, -18,  -7,  -1,  4,  -9, -15, -27,
	-23,  -9, -23,  -5, -9, -16,  -5, -17,
})

var earlyRookTable [64]int = reverse(447,[64]int{
	 32,  42,  32,  51, 63,  9,  31,  43,
	 27,  32,  58,  62, 80, 67,  26,  44,
	 -5,  19,  26,  36, 17, 45,  61,  16,
	-24, -11,   7,  26, 24, 35,  -8, -20,
	-36, -26, -12,  -1,  9, -7,   6, -23,
	-45, -25, -16, -17,  3,  0,  -5, -33,
	-44, -16, -20,  -9, -1, 11,  -6, -71,
	-19, -13,  -5,  15, 12,  7, -37, -26,
})

var lateRookTable [64]int = reverse(512,[64]int{
	13, 10, 18, 15, 12,  12,   8,   5,
	11, 13, 13, 11, -3,   3,   8,   3,
	 7,  7,  7,  5,  4,  -3,  -5,  -3,
	 4,  3, 13,  1,  2,   1,  -1,   2,
	 3,  5,  8,  4, -5,  -6,  -8, -11,
	-4,  0, -5, -1, -7, -12,  -8, -16,
	-6, -6,  0,  2, -9,  -9, -11,  -3,
	-9,  2,  3, -1, -5, -13,   4, -20,
})


var earlyQueenTable [64]int = reverse(1025,[64]int{
	-28,   0,  29,  12,  59,  44,  43,  45,
	-24, -39,  -5,   1, -16,  57,  28,  54,
	-13, -17,   7,   8,  29,  56,  47,  57,
	-27, -27, -16, -16,  -1,  17,  -2,   1,
	 -9, -26,  -9, -10,  -2,  -4,   3,  -3,
	-14,   2, -11,  -2,  -5,   2,  14,   5,
	-35,  -8,  11,   2,   8,  15,  -3,   1,
	 -1, -18,  -9,  -5, -15, -25, -31, -50,
})

var lateQueenTable [64]int = reverse(936,[64]int{
	 -9,  22,  22,  27,  27,  19,  10,  20,
	-17,  20,  32,  41,  58,  25,  30,   0,
	-20,   6,   9,  49,  47,  35,  19,   9,
	  3,  22,  24,  45,  57,  40,  57,  36,
	-18,  28,  19,  47,  31,  34,  39,  23,
	-16, -27,  15,   6,   9,  17,  10,   5,
	-22, -23, -30, -16, -16, -23, -36, -32,
	-33, -28, -22, -43,  -5, -32, -20, -41,
})

var earlyKingTable [64]int = reverse(0,[64]int{
	-65,  23,  16, -15, -56, -34,   2,  13,
	 29,  -1, -20,  -7,  -8,  -4, -38, -29,
	 -9,  24,   2, -16, -20,   6,  22, -22,
	-17, -20, -12, -27, -30, -25, -14, -36,
	-49,  -1, -27, -39, -46, -44, -33, -51,
	-14, -14, -22, -46, -44, -30, -15, -27,
	  1,   7,  -8, -64, -43, -16,   9,   8,
	-15,  36,  12, -54,   8, -28,  30,  14,
})

var lateKingTable [64]int = reverse(0,[64]int{
	-74, -35, -18, -18, -11,  15,   4, -17,
	-12,  17,  14,  17,  17,  38,  23,  11,
	 10,  17,  23,  15,  20,  45,  44,  13,
	 -8,  22,  24,  27,  26,  33,  26,   3,
	-18,  -4,  21,  24,  27,  23,   9, -11,
	-19,  -3,  11,  21,  23,  16,   7,  -9,
	-27, -11,   4,  13,  14,   4,  -5, -17,
	-53, -34, -21, -11, -28, -14, -24, -43,
})

func CustomV1(board *dragontoothmg.Board) float64 {
	phase := 0
	midScore := 0
	endScore := 0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				midScore += earlyPawnTable[i]
				endScore += latePawnTable[i]
			} else if board.White.Knights>>i%2 == 1 {
				midScore += earlyKnightTable[i]
				endScore += lateKnightTable[i]
				phase += 1
			} else if board.White.Bishops>>i%2 == 1 {
				midScore += earlyBishopTable[i]
				endScore += lateBishopTable[i]
				phase += 1
			} else if board.White.Rooks>>i%2 == 1 {
				midScore += earlyRookTable[i]
				endScore += lateRookTable[i]
				phase += 2
			} else if board.White.Queens>>i%2 == 1 {
				midScore += earlyQueenTable[i]
				endScore += lateQueenTable[i]
				phase += 4
			} else if board.White.Kings>>i%2 == 1{
				midScore += earlyKingTable[i]
				endScore += lateKingTable[i]
			}
		} else if board.Black.All>>i%2 == 1 {
			j := i ^ 56
			if board.Black.Pawns>>i%2 == 1 {
				midScore -= earlyPawnTable[j]
				endScore -= latePawnTable[j]
			} else if board.Black.Knights>>i%2 == 1 {
				midScore -= earlyKnightTable[j]
				endScore -= lateKnightTable[j]
				phase += 1
			} else if board.Black.Bishops>>i%2 == 1 {
				midScore -= earlyBishopTable[j]
				endScore -= lateBishopTable[j]
				phase += 1
			} else if board.Black.Rooks>>i%2 == 1 {
				midScore -= earlyRookTable[j]
				endScore -= lateRookTable[j]
				phase += 2
			} else if board.Black.Queens>>i%2 == 1 {
				midScore -= earlyQueenTable[j]
				endScore -= lateQueenTable[j]
				phase += 4
			} else if board.Black.Kings>>i%2 == 1{
				midScore -= earlyKingTable[j]
				endScore -= lateKingTable[j]
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

func CustomV2(board *dragontoothmg.Board) float64 {
	phase := 0
	midScore := 0
	endScore := 0

	for i := 0; i < 64; i++ {
		if board.White.All>>i%2 == 1 {
			if board.White.Pawns>>i%2 == 1 {
				midScore += earlyPawnTable[i]
				endScore += latePawnTable[i]
			} else if board.White.Knights>>i%2 == 1 {
				midScore += earlyKnightTable[i]
				endScore += lateKnightTable[i]
				phase += 1
			} else if board.White.Bishops>>i%2 == 1 {
				midScore += earlyBishopTable[i]
				endScore += lateBishopTable[i]
				phase += 1
			} else if board.White.Rooks>>i%2 == 1 {
				midScore += earlyRookTable[i]
				endScore += lateRookTable[i]
				phase += 2
			} else if board.White.Queens>>i%2 == 1 {
				midScore += earlyQueenTable[i]
				endScore += lateQueenTable[i]
				phase += 4
			} else if board.White.Kings>>i%2 == 1{
				midScore += earlyKingTable[i]
				endScore += lateKingTable[i]
			}
		} else if board.Black.All>>i%2 == 1 {
			j := i ^ 56
			if board.Black.Pawns>>i%2 == 1 {
				midScore -= earlyPawnTable[j]
				endScore -= latePawnTable[j]
			} else if board.Black.Knights>>i%2 == 1 {
				midScore -= earlyKnightTable[j]
				endScore -= lateKnightTable[j]
				phase += 1
			} else if board.Black.Bishops>>i%2 == 1 {
				midScore -= earlyBishopTable[j]
				endScore -= lateBishopTable[j]
				phase += 1
			} else if board.Black.Rooks>>i%2 == 1 {
				midScore -= earlyRookTable[j]
				endScore -= lateRookTable[j]
				phase += 2
			} else if board.Black.Queens>>i%2 == 1 {
				midScore -= earlyQueenTable[j]
				endScore -= lateQueenTable[j]
				phase += 4
			} else if board.Black.Kings>>i%2 == 1{
				midScore -= earlyKingTable[j]
				endScore -= lateKingTable[j]
			}
		}
	}


	if phase > 24 {
		phase = 24
	}

	eval := float64((midScore * phase) + (endScore * (24-phase)))/24/100 

	if board.OurKingInCheck() {
		if board.Wtomove {
			eval -= 1.5
		} else {
			eval += 1.5
		}
	}


	if !board.Wtomove {
		eval = -eval
	}

	

	
	
	// An idea taken from stockfish called "contempt"
	// It tries to encourage simplifying positions while ahead in material and avoiding simplifications when behind
	const discontempt float64 = 0.2
	eval *= 1 + (discontempt * float64(24 - phase) / 24)

	return SigmoidLike(eval)
}