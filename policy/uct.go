package policy

import (
	"github.com/ffloof/dragontoothmg"
)

func UCT(board *dragontoothmg.Board, move dragontoothmg.Move, threats *[64]int8) float64 {
	return 1.0
}

//nothing, pawn, knight, bishop, rook, queen, king

// Tables are in format of x takes y
var LossOdds = [5][6]float64{
	{2.4, 2.8, 3.8, 3.8, 4.5, 5.6}, // Pawn takes Y
	{2.0, 2.1, 2.8, 2.8, 3.1, 3.8}, // Knight takes Y
	{2.0, 2.1, 2.8, 2.8, 3.1, 3.8}, // Bishop takes Y
	{1.8, 2.0, 2.5, 2.5, 2.8, 3.4}, // Rook takes Y
	{1.7, 1.8, 2.1, 2.1, 2.4, 2.8}, // Queen takes Y
	// King takes Y isnt possible as the piece is defended
}

var WinOdds = [6][6]float64{
	{2.8, 3.4, 4.2, 4.2, 4.8, 5.8}, // Pawn takes Y
	{2.8, 3.0, 3.4, 3.4, 3.7, 4.2}, // Knight takes Y
	{2.8, 3.0, 3.4, 3.4, 3.7, 3.7}, // Bishop takes Y
	{2.8, 3.0, 3.3, 3.3, 3.4, 3.7}, // Rook takes Y
	{2.8, 3.0, 2.1, 2.1, 3.2, 3.4}, // Queen takes Y
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