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
	MinMax   float64
	Expanded bool
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
		history = history
		
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
			node.Children = []MonteCarloNode{}
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
		bestMinMax := -1.0
		parentConstant := math.Log(node.Visits) // Exploration parameter is now factored into HeurUCT

		for i := range node.Moves {
			var score float64
			if i >= len(node.Children) {
				const discovery float64 = 0.5
				score = mcts.treeFunc(board, node.Moves[i], node.Threats) * math.Sqrt(parentConstant/discovery)
			} else {
				child := &node.Children[i]
				average := -child.Value / child.Visits
				score = average + (mcts.treeFunc(board, node.Moves[i], node.Threats) * math.Sqrt(parentConstant/child.Visits))

				if -child.MinMax > bestMinMax {
					bestMinMax = -child.MinMax
				}
			}
			
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

		if len(node.Children) != 0 {
			node.MinMax = bestMinMax
		}

		board.Apply(node.Moves[bestChildIndex])
		
		j := len(node.Children)

		if bestChildIndex >= j {
			evaluation = mcts.evalFunc(board)
			node.Children = append(node.Children, MonteCarloNode{Parent: node, Value:evaluation, Visits:1, MinMax:clamp1(evaluation)})
			node.Moves[bestChildIndex], node.Moves[j] = node.Moves[j], node.Moves[bestChildIndex]

			node = &node.Children[j]

			break selectionLoop
		} else {
			nextNode := &node.Children[bestChildIndex]
			if nextNode.Parent != node { // Should avoid slice reallocations breaking stuff
				nextNode.Parent = node
			}
			node = nextNode
		}
	}

	// 3. Backpropogation

	eliminated := false
	node = node.Parent
	for node != nil {
		evaluation = -evaluation
		node.Visits++

		if !eliminated {
			if node.MinMax < evaluation {
				node.MinMax = clamp1(evaluation)
				node.Value += evaluation
			} else {
				eliminated = true
				node.Value += (evaluation + node.MinMax) / 2
			}
		} else {
			node.Value += (evaluation + node.MinMax) / 2
		}

		
		node = node.Parent
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

func clamp1(n float64) float64 {
	if n > 1 {
		return 1
	} else if n < -1 {
		return -1
	}
	return n
}