package mcts

import (
	"github.com/dylhunn/dragontoothmg"
)

type MonteCarloTreeSearcher struct {
	BaseState dragontoothmg.Board
	Head *MonteCarloNode
	treeFunc func(*MonteCarloNode ,*MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64
	evalFunc func(dragontoothmg.Board) float64
}

func NewSearch(board dragontoothmg.Board, tree func(*MonteCarloNode ,*MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64, eval func(dragontoothmg.Board) float64) MonteCarloTreeSearcher {
	return MonteCarloTreeSearcher {
		BaseState : board,
		Head: newNode(nil, board),
		treeFunc:tree,
		evalFunc:eval,
	}
}

func (mcts MonteCarloTreeSearcher) RunIterations(n int){
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

func (mcts *MonteCarloTreeSearcher) ApplyMove(move dragontoothmg.Move){
	mcts.BaseState.Apply(move)
	for i, option := range mcts.Head.Moves {
		if option == move {
			selectedChild := mcts.Head.Children[i]
			if selectedChild == nil {
				break
			}

			mcts.Head = selectedChild
			mcts.Head.Parent = nil
			return
		}
	}

	mcts.Head = newNode(nil, mcts.BaseState)
}

func (mcts *MonteCarloTreeSearcher) ApplyMoveString(movestr string) error {
	move, err := dragontoothmg.ParseMove(movestr)
	if err != nil {
		return err 
	}
	
	mcts.ApplyMove(move)
	return nil
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
		// If any null node exists expand it otherwise choose the one with best uct score
		
		for i := range(node.Children) {
			if node.Children[i] == nil {
				// 2. Expansion and Evaluation
				board.Apply(node.Moves[i])
				
				nextNode := newNode(node, board) 
				node.Children[i] = nextNode
				
				node = nextNode
				evaluation = mcts.evalFunc(board)
				break selectionLoop
			}
		}

		bestChildIndex := 0
		bestScore := -1.0
		for i, v := range(node.Children) {
			score := mcts.treeFunc(node, v, board, node.Moves[i])
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

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

func newNode(parent *MonteCarloNode, board dragontoothmg.Board) *MonteCarloNode{
	moves := board.GenerateLegalMoves()

	children := make([]*MonteCarloNode, len(moves), len(moves))

	return &MonteCarloNode{
		Parent: parent,
		Children: children,
		Moves: moves,
		Value: 0.0,
		Visits: 0.0,
	}
}

type MonteCarloNode struct {
	Parent *MonteCarloNode
	Children []*MonteCarloNode //TODO: consider making this just a list of nodes and run benchmarks
	Moves []dragontoothmg.Move
	Value float64
	Visits float64
}