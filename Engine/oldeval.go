package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func CountMaterial(board *dragontoothmg.Board) int16 {
	var count int16

	for i := 0; i < 64; i++ {
		if board.White.All >> i % 2 == 1 {
			if board.White.Pawns >> i % 2 == 1 {
				count += 100
			} else if board.White.Knights >> i % 2 == 1 {
				count += 280
			} else if board.White.Bishops >> i % 2 == 1 {
				count += 320
			} else if board.White.Rooks >> i % 2 == 1 {
				count += 500
			} else if board.White.Queens >> i % 2 == 1 {
				count += 900
			}
		} else if board.Black.All >> i % 2 == 1 {
			if board.Black.Pawns >> i % 2 == 1 {
				count -= 100
			} else if board.Black.Knights >> i % 2 == 1 {
				count -= 280
			} else if board.Black.Bishops >> i % 2 == 1 {
				count -= 320
			} else if board.Black.Rooks >> i % 2 == 1 {
				count -= 500
			} else if board.Black.Queens >> i % 2 == 1 {
				count -= 900
			}
		}
	}	

	if !board.Wtomove {
		count = -count
	}

	return count
}