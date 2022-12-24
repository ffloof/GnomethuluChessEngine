package search

import (
	"github.com/ffloof/dragontoothmg"
	"math"
)

type MonteCarloNode struct {
	Parent   *MonteCarloNode
	Children []MonteCarloNode
	Moves    []dragontoothmg.Move
	Value    float64
	Visits   float64
	Threats *[64]int8
	Expanded bool
}

func newNode(parent *MonteCarloNode) MonteCarloNode {
	if parent != nil {
		return MonteCarloNode{
			Parent:   parent,
		}
	} else {
		return MonteCarloNode{
			Parent: parent,
			Children: []MonteCarloNode{},
		}
	}
}



// Main search functionality here
func (mcts *MonteCarloTreeSearcher) iteration() {
	history := map[uint64]bool{}

	evaluation := 0.0

	// 1. Selection
	copiedBoard := mcts.startPos
	board := &copiedBoard
	node := mcts.Head

selectionLoop:
	for true {
		if history[board.Hash()] {
			evaluation = 0.0
			break selectionLoop
		} else {
			history[board.Hash()] = true
		}

		//TODO: probably merge threats and expansion unless we use them in eval
		if node.Threats == nil {
			node.Threats = ControlMap(board)
		}

		if !node.Expanded {
			node.Expanded = true
			node.Moves = board.GenerateLegalMoves()
		}

		if len(node.Moves) == 0 {
			if board.OurKingInCheck() {
				evaluation = -MateAdjust(node)
			} else {
				evaluation = 0.0
			}
			break selectionLoop
		}

		bestChildIndex := 0
		bestScore := -1.0
		parentConstant := mcts.PolicyExplore * math.Log(node.Visits)

		for i := range node.Moves {
			var score float64
			if i >= len(node.Children) {
				const discovery float64 = 100
				score = mcts.treeFunc(board, node.Moves[i], node.Threats) * discovery
			} else {
				child := &node.Children[i]
				score = (-child.Value / child.Visits) + (mcts.treeFunc(board, node.Moves[i], node.Threats) * math.Sqrt(parentConstant/child.Visits))
			}
			
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

		board.Apply(node.Moves[bestChildIndex])
		
		j := len(node.Children)
		if bestChildIndex >= j {
			node.Children = append(node.Children, newNode(node))
			node.Moves[bestChildIndex], node.Moves[j] = node.Moves[j], node.Moves[bestChildIndex]
			node = &node.Children[j]
			evaluation = mcts.evalFunc(board)

			break selectionLoop
		} else {
			node = &node.Children[bestChildIndex]
		}
	}

	// 3. Backpropogation

	for node != nil {
		node.Visits++
		node.Value += evaluation
		node = node.Parent
		evaluation = -evaluation
	}
}



// When a mate in 1 since it is a terminal state win any previously explored branches of the parent are irrelevant, since it will always opt for mate in 1
// So MateAdjust() will return the eval backpropogate to correct this difference and remove all other branches from parent node
// This helps the algorithm find mates exponentially faster than it otherwise would
func MateAdjust(node *MonteCarloNode) float64 {
	for i := range node.Parent.Children {
		child := &node.Parent.Children[i]
		if node == child {
			// Make the mate the only possible node
			move := node.Parent.Moves[i]
			node.Parent.Children = []MonteCarloNode{*node}
			node.Parent.Moves = []dragontoothmg.Move{move}
			break

			// The value of the parent should be equal to visits since it has won the game
			// So backpropogation will propogate node.Parent.Visits - node.Parent.Value to correct the value up the tree
			// (node.Parent.Visits - node.Parent.Value) + node.Parent.Value = node.Parent.Visits	
		}
	}
	return node.Parent.Visits - node.Parent.Value
}