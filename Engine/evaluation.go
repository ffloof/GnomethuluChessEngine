package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

const pawn int16 = 100
const knight int16 = 280
const bishop int16 = 320
const rook int16 = 500
const queen int16 = 900

func CountMaterial(board *dragontoothmg.Board) int16 {
	var count int16

	for i := 0; i < 64; i++ {
		if board.White.All >> i % 2 == 1 {
			if board.White.Pawns >> i % 2 == 1 {
				count += pawn
			} else if board.White.Knights >> i % 2 == 1 {
				count += knight
			} else if board.White.Bishops >> i % 2 == 1 {
				count += bishop
			} else if board.White.Rooks >> i % 2 == 1 {
				count += rook
			} else if board.White.Queens >> i % 2 == 1 {
				count += queen
			}
		} else if board.Black.All >> i % 2 == 1 {
			if board.Black.Pawns >> i % 2 == 1 {
				count -= pawn
			} else if board.Black.Knights >> i % 2 == 1 {
				count -= knight
			} else if board.Black.Bishops >> i % 2 == 1 {
				count -= bishop
			} else if board.Black.Rooks >> i % 2 == 1 {
				count -= rook
			} else if board.Black.Queens >> i % 2 == 1 {
				count -= queen
			}
		}
	}	

	if !board.Wtomove {
		count = -count
	}

	return count
}