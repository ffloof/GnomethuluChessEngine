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
	{1.7, 2.0, 2.7, 2.7, 3.2, 4.0}, // Pawn takes Y
	{1.4, 1.5, 2.0, 2.0, 2.2, 2.7}, // Knight takes Y
	{1.4, 1.5, 2.0, 2.0, 2.2, 2.7}, // Bishop takes Y
	{1.3, 1.4, 1.8, 1.8, 2.0, 2.4}, // Rook takes Y
	{1.2, 1.3, 1.5, 1.5, 1.7, 2.0}, // Queen takes Y
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, // King takes Y (this shouldnt be possible as king should never be able to take a defended piece)
}

var WinOdds = [6][6]float64{
	{2.0, 2.4, 3.0, 3.0, 3.4, 4.1}, // Pawn takes Y
	{2.0, 2.1, 2.4, 2.4, 2.6, 3.0}, // Knight takes Y
	{2.0, 2.1, 2.4, 2.4, 2.6, 3.0}, // Bishop takes Y
	{2.0, 2.1, 2.3, 2.3, 2.4, 2.6}, // Rook takes Y
	{2.0, 2.1, 2.2, 2.2, 2.3, 2.4}, // Queen takes Y
	{2.0, 2.0, 2.0, 2.0, 2.0, 2.0}, // King takes Y
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