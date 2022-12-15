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
	{0.7, 1.0, 1.7, 1.7, 2.2, 3.0}, // Pawn takes Y
	{0.4, 0.5, 1.0, 1.0, 1.2, 1.7}, // Knight takes Y
	{0.4, 0.5, 1.0, 1.0, 1.2, 1.7}, // Bishop takes Y
	{0.3, 0.4, 0.8, 0.8, 1.0, 1.4}, // Rook takes Y
	{0.2, 0.3, 0.5, 0.5, 0.7, 1.0}, // Queen takes Y
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, // King takes Y (this shouldnt be possible as king should never be able to take a defended piece)
}

var WinOdds = [6][6]float64{
	{1.0, 1.4, 2.0, 2.0, 2.4, 3.1}, // Pawn takes Y
	{1.0, 1.1, 1.4, 1.4, 1.6, 2.0}, // Knight takes Y
	{1.0, 1.1, 1.4, 1.4, 1.6, 2.0}, // Bishop takes Y
	{1.0, 1.1, 1.3, 1.3, 1.4, 1.6}, // Rook takes Y
	{1.0, 1.1, 1.2, 1.2, 1.3, 1.4}, // Queen takes Y
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, // King takes Y
}

func HeurUCT(board *dragontoothmg.Board, move dragontoothmg.Move, threats *[64]int8) float64 {
	aggresor, _ := dragontoothmg.GetPieceType(move.From(), board)
	victim, _ := dragontoothmg.GetPieceType(move.To(), board)
	//TODO: implement another factor "risk", essentially if the square the piece is standing on is in danger, we need to move

	var odds float64
	//TODO: implement castling bonus
	if threats[move.To()] <= 0 {
		odds = LossOdds[aggresor - 1][victim]
	} else {
		odds = WinOdds[aggresor - 1][victim]
	}

	if threats[move.From()] < 0 {
		odds *= 1.2
	}
	return odds
}