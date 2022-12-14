package policy

import (
	"github.com/ffloof/dragontoothmg"
)

func UCT(board *dragontoothmg.Board, move dragontoothmg.Move, threats *[64]int8) float64 {
	return 1.0
}

//nothing, pawn, knight, bishop, rook, queen, king

// Tables are in format of x takes y
var LossOdds = [6][6]float64{
	{0.5, 1.0, 3.0, 3.0, 5.0, 9.0}, // Pawn takes Y
	{0.2, 0.3, 1.0, 1.0, 1.5, 3.0}, // Knight takes Y
	{0.2, 0.3, 1.0, 1.0, 1.5, 3.0}, // Bishop takes Y
	{0.1, 0.2, 1.0, 1.0, 1.0, 2.0}, // Rook takes Y
	{0.1, 0.1, 0.3, 0.3, 0.5, 1.0}, // Queen takes Y
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, // King takes Y (this shouldnt be possible as king should never be able to take a defended piece)
}

var WinOdds = [6][6]float64{
	{1.0, 2.0, 4.0, 4.0, 6.0, 10.0}, // Pawn takes Y
	{1.0, 1.3, 2.0, 2.0, 2.7, 4.0}, // Knight takes Y
	{1.0, 1.3, 2.0, 2.0, 2.7, 4.0}, // Bishop takes Y
	{1.0, 1.2, 1.6, 1.6, 2.0, 2.8}, // Rook takes Y
	{1.0, 1.1, 1.3, 1.3, 1.5, 2.0}, // Queen takes Y
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, // King takes Y
}

func HeurUCT(board *dragontoothmg.Board, move dragontoothmg.Move, threats *[64]int8) float64 {
	aggresor, _ := dragontoothmg.GetPieceType(move.From(), board)
	victim, _ := dragontoothmg.GetPieceType(move.To(), board)

	//TODO: implement castling bonus
	if threats[move.To()] < 0 {
		return LossOdds[aggresor - 1][victim]
	} else {
		return WinOdds[aggresor - 1][victim]
	}
	
}