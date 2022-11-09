package mcts

import (
	"fmt"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

type MonteCarloTreeSearcher struct {
	BaseState dragontoothmg.Board
	Head *monteCarloNode
	//TODO: add eval func and tree policy func
}

func NewSearch(board dragontoothmg.Board) MonteCarloTreeSearcher {
	return MonteCarloTreeSearcher {
		BaseState : board,
		Head: newNode(nil, board, []dragontoothmg.Move{}),
	}
}


func (mcts MonteCarloTreeSearcher) RunIterations(n int){
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

func (mcts MonteCarloTreeSearcher) GetBestMove() dragontoothmg.Move {
	bestIndex := 0
	bestAverage := -1.0
	for i, v := range mcts.Head.Children {
		if v.Value / v.Visits > bestAverage {
			bestIndex = i
			bestAverage = v.Value / v.Visits
		}
	}
	return mcts.Head.Moves[bestIndex]
}

func (mcts MonteCarloTreeSearcher) iteration(){
	TEMPSTACK := []dragontoothmg.Move{}
	evaluation := 0.0

	// 1. Selection
	board := mcts.BaseState
	node := mcts.Head


	selectionLoop:
	for true {
		if len(node.Moves) == 0 {
			if board.OurKingInCheck() {
				evaluation = 1.0
			} else {
				evaluation = 0.0
			}
			break selectionLoop	
		}
		//TODO: deal with checking if the game is won
		// If any null node exists expand it otherwise choose the one with best uct score
		
		for i := range(node.Children) {
			if node.Children[i] == nil {
				// 2. Expansion and Evaluation
				board.Apply(node.Moves[i])
				TEMPSTACK = append(TEMPSTACK, node.Moves[i])
				
				nextNode := newNode(node, board, TEMPSTACK) 
				node.Children[i] = nextNode
				
				node = nextNode
				evaluation = Evaluate(board)
				break selectionLoop
			}
		}

		bestChildIndex := 0
		bestScore := -1.0
		for i, v := range(node.Children) {
			score := UCT(node, v)
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

		
		TEMPSTACK = append(TEMPSTACK, node.Moves[bestChildIndex])
		board.Apply(node.Moves[bestChildIndex])
		node = node.Children[bestChildIndex]
	} 

	// 3. Backpropogation
	

	for node != nil {
		node.Visits++
		node.Value += evaluation

		evaluation = -evaluation
		node = node.Parent
	}
}

func newNode(parent *monteCarloNode, board dragontoothmg.Board, TEMPSTACK []dragontoothmg.Move) *monteCarloNode{
	defer func() {
		if err := recover(); err != nil {

			fmt.Println(board.ToFen())
			fmt.Println(TEMPSTACK)
			for _, v := range TEMPSTACK {
				fmt.Println(v.String())
			}

			panic("new panic 😎")
			//fmt.Println("panic occurred:", err)
		}
	}()
	moves := board.GenerateLegalMoves()

	children := make([]*monteCarloNode, len(moves), len(moves))

	return &monteCarloNode{
		Parent: parent,
		Children: children,
		Moves: moves,
		Value: 0.0,
		Visits: 0.0,
	}
}

func UCT(parent, child *monteCarloNode) float64 {
	c := 0.4
	return (child.Value/child.Visits) + math.Sqrt(c*math.Log(parent.Visits)/child.Visits) 
}

func Evaluate(board dragontoothmg.Board) float64 {
	eval := 0.0

	for i := 0; i <64 ;i++ {
		if board.White.All >> i % 2 == 1 {
			if board.White.Pawns >> i % 2 == 1 {
				eval += 1.0
			}
			if board.White.Knights >> i % 2 == 1 {
				eval += 3.0
			}
			if board.White.Bishops >> i % 2 == 1 {
				eval += 3.0
			}
			if board.White.Rooks >> i % 2 == 1 {
				eval += 5.0
			}
			if board.White.Queens >> i % 2 == 1 {
				eval += 9.0
			}
		}
		if board.Black.All >> i % 2 == 1 {
			if board.Black.Pawns >> i % 2 == 1 {
				eval -= 1.0
			}
			if board.Black.Knights >> i % 2 == 1 {
				eval -= 3.0
			}
			if board.Black.Bishops >> i % 2 == 1 {
				eval -= 3.0
			}
			if board.Black.Rooks >> i % 2 == 1 {
				eval -= 5.0
			}
			if board.Black.Queens >> i % 2 == 1 {
				eval -= 9.0
			}
		}
	}

	eval/=15
	// TODO: make sure this doesnt lead to missing mate

	if board.Wtomove {
		eval = -eval
	}

	return eval
}

type monteCarloNode struct {
	Parent *monteCarloNode
	Children []*monteCarloNode
	Moves []dragontoothmg.Move
	Value float64
	Visits float64
}