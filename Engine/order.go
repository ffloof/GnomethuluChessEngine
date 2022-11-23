package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

//TODO: create another implementation where it just simply sorts moves, benchmark accordingly
type moveOrder struct {
	moves []dragontoothmg.Move
	priority map[dragontoothmg.Move]int
	startIndex int
}

func CreateMoveOrder(board *dragontoothmg.Board, tableMove dragontoothmg.Move, inQuiescenceSearch bool) *moveOrder {
	moves := board.GenerateLegalMoves()
	priority := map[dragontoothmg.Move]int{}
	
	for _, move := range moves {
		priority[move] = getPriority(board, move, tableMove, inQuiescenceSearch)
	}

	return &moveOrder {
		moves: moves,
		priority: priority,
		startIndex: 0,
	}
}

func (order *moveOrder) NoMoves() bool {
	return len(order.moves) == 0
}

func (order *moveOrder) GetNextMove() (dragontoothmg.Move,bool) {
	bestIndex := -1
	bestPriority := -1

	for i := order.startIndex; i < len(order.moves); i++ {
		if order.priority[order.moves[i]] > bestPriority {
			bestIndex = i
			bestPriority = order.priority[order.moves[i]] 
		}
	}

	if bestIndex < 0 {
		order.startIndex = len(order.moves)
		return 0, true
	}
	
	chosenMove := order.moves[bestIndex]
	order.moves[order.startIndex], order.moves[bestIndex] = order.moves[bestIndex], order.moves[order.startIndex]
	order.startIndex++
	return chosenMove, false
}



var MVVLVAPriority = [6][7]int {
	{0,0,0,0,0,0,0}, //Nothing taken by X
	{16,15,14,13,12,11,10}, //Pawn taken by X
	{26,25,24,23,22,21,20}, //Knight taken by X
	{27,26,25,24,23,22,21}, //Bishop taken by X
	{36,35,34,33,32,31,30}, //Rook taken by X
	{46,45,44,43,42,41,40}, //Queen taken by X
}

//TODO: deal with problem where a check is caused in quiescence search, its fine for now but might want to be addressed later
func getPriority (board *dragontoothmg.Board, move dragontoothmg.Move, tableMove dragontoothmg.Move, inQuiescenceSearch bool) int {
	if move == tableMove {
		return 1000
	} else if piece := move.Promote(); piece != dragontoothmg.Nothing {
		if piece == dragontoothmg.Queen {
			return 100
		} 
		return -1
	} else if dragontoothmg.IsCapture(move, board) {
		victim, _ := dragontoothmg.GetPieceType(move.To(), board)
		attacker, _ := dragontoothmg.GetPieceType(move.From(), board)
		return MVVLVAPriority[victim][attacker]
	} else {
		if inQuiescenceSearch {
			return -1
		} else {
			return 0
		}
	}
}

