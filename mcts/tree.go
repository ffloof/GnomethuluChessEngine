package mcts

import (
	"github.com/dylhunn/dragontoothmg"
)

func (mcts MonteCarloTreeSearcher) iteration() {
	evaluation := 0.0

	// 1. Selection
	board := mcts.BaseState
	node := mcts.Head

	previousBoards := map[dragontoothmg.Board]bool{}

selectionLoop:
	for true {
		if previousBoards[board] {
			evaluation = 0.0
			break selectionLoop
		}

		if len(node.Moves) == 0 {
			if board.OurKingInCheck() {
				evaluation = 1.0
			} else {
				evaluation = 0.0
			}
			break selectionLoop
		}
		// If any null node exists expand it otherwise choose the one with best uct score

		for i := range node.Children {
			if node.Children[i] == nil {
				// 2. Expansion and Evaluation
				previousBoards[board] = true
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
		for i, v := range node.Children {
			score := mcts.treeFunc(node, v, board, node.Moves[i])
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

		previousBoards[board] = true
		board.Apply(node.Moves[bestChildIndex])
		node = node.Children[bestChildIndex]
	}

	// 3. Backpropogation

	endMinMax := false
	for node != nil {
		node.Visits++
		node.Value += evaluation
		if !endMinMax { //TODO: this minmax doesnt work perfectly
			if evaluation > node.MinMax {
				node.MinMax = evaluation
			} else {
				endMinMax = true
			}
		}

		evaluation = -evaluation
		node = node.Parent
	}
}

func newNode(parent *MonteCarloNode, board dragontoothmg.Board) *MonteCarloNode {
	moves := board.GenerateLegalMoves()

	children := make([]*MonteCarloNode, len(moves), len(moves))

	return &MonteCarloNode{
		Parent:   parent,
		Children: children,
		Moves:    moves,
		Value:    0.0,
		Visits:   0.0,
		MinMax:   -1.0,
	}
}

type MonteCarloNode struct {
	Parent   *MonteCarloNode
	Children []*MonteCarloNode //TODO: consider making this just a list of nodes and run benchmarks
	Moves    []dragontoothmg.Move
	Value    float64
	Visits   float64
	MinMax   float64
}
