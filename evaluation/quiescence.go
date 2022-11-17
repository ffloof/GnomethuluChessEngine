package evaluation

import (
	"sort"
	"github.com/dylhunn/dragontoothmg"
)

// Consider transposition table and adding checks
func pestoQuiescence(board dragontoothmg.Board, alpha, beta float64) float64 {
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
					capture_moves = append(capture_moves, move)
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

		sort.Slice(capture_moves, Less_MVV_LVA)

		chosen_moves = append(promote_moves, capture_moves...)
		//chosen_moves = capture_moves
	}

	for _, move := range chosen_moves {
		if dragontoothmg.IsCapture(move, &board) {
			undo := board.Apply(move) 
			
			score = -pestoQuiescence(board, -beta, -alpha)
			
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


func pestoQuiescenceChecks(board dragontoothmg.Board, alpha, beta float64) float64 {
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
	check_moves := []dragontoothmg.Move{}

	if board.OurKingInCheck() {
		chosen_moves = all_moves
	} else {
		for _, move := range all_moves {
			promotePiece := move.Promote()
			if promotePiece == dragontoothmg.Nothing {
				if dragontoothmg.IsCapture(move, &board) {
					capture_moves = append(capture_moves, move)
				} 
			} else if promotePiece == dragontoothmg.Queen { //Queen
				promote_moves = append(promote_moves, move)
			} else {
				undo := board.Apply(move)
				if board.OurKingInCheck() {
					check_moves = append(check_moves, move)
				}
				undo()
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

		sort.Slice(capture_moves, Less_MVV_LVA)

		chosen_moves = append(append(promote_moves, check_moves...), capture_moves...)
	}

	for _, move := range chosen_moves {
		if dragontoothmg.IsCapture(move, &board) {
			undo := board.Apply(move) 
			
			score = -pestoQuiescence(board, -beta, -alpha)
			
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