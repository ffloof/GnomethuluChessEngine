package policy

import (
	"github.com/dylhunn/dragontoothmg"
)

func UCT(board *dragontoothmg.Board, move dragontoothmg.Move) float64 {
	return 1.0
}

var MVVLVA = [6][6]float64{
//X = -    P    N    B    R    Q  
	{1.1, 1.0, 3.0, 3.0, 5.0, 9.0}, //Pawn takes X
	{1.0, 1.0, 1.0, 1.0, 1.5, 3.0}, //Knight takes X
	{1.0, 1.0, 1.0, 1.0, 1.5, 3.0}, //Bishop takes X
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.8}, //Rook takes X
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, //Queen takes X
	{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}, //King takes X
}

func MVVLVA_UCT(board *dragontoothmg.Board, move dragontoothmg.Move) float64 {
	attacker, _ := dragontoothmg.GetPieceType(move.From(), board)
	victim, _ := dragontoothmg.GetPieceType(move.To(), board)
	return MVVLVA[attacker-1][victim]
}