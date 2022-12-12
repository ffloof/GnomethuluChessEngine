package search

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
)

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

		if len(node.Moves) == 0 {
			if !node.FullyExpanded {
				node.Moves = board.GenerateLegalMoves()
				node.Children = make([]*MonteCarloNode, len(node.Moves), len(node.Moves))
			}

			if len(node.Moves) == 0 {
				if board.OurKingInCheck() {
					evaluation = -MateAdjust(node)
				} else {
					evaluation = 0.0
				}
				break selectionLoop
			}
		}
		
		// If any null node exists expand it otherwise choose the one with best uct score
		if !node.FullyExpanded {
			nowFull := true
			for i := range node.Children {
				if node.Children[i] == nil {
					// 2. Expansion and Evaluation
					board.Apply(node.Moves[i])

					nextNode := newNode(node, board)
					node.Children[i] = nextNode

					node = nextNode
					evaluation = mcts.evalFunc(board)
					nowFull = false
					break selectionLoop
				}
			}

			if nowFull {
				node.FullyExpanded = true
			}
		}

		bestChildIndex := 0
		bestScore := -1.0
		parentConstant := mcts.PolicyExplore * math.Log(node.Visits)
		for i, child := range node.Children {
			//score := mcts.treeFunc(parentConstant, child, board, node.Moves[i])
			score := (-child.Value / child.Visits) + (mcts.treeFunc(board, node.Moves[i]) * math.Sqrt(parentConstant/child.Visits))
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
		node = node.Parent
		evaluation = -evaluation
	}
}

func newNode(parent *MonteCarloNode, board *dragontoothmg.Board) *MonteCarloNode {
	return &MonteCarloNode{
		Parent:   parent,
	}
}

type MonteCarloNode struct { //TODO: look into if we can garbage collect some nodes or at least node.Moves
	Parent   *MonteCarloNode
	Children []*MonteCarloNode //TODO: try to mitigate pointer chasing
	Moves    []dragontoothmg.Move
	Value    float64
	Visits   float64
	FullyExpanded bool
}

func NewSearch(tree func(*dragontoothmg.Board, dragontoothmg.Move) float64, eval func(*dragontoothmg.Board) float64) MonteCarloTreeSearcher {
	board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	mcts := MonteCarloTreeSearcher{
		startPos: board,
		Head:     newNode(nil, &board),
		treeFunc: tree,
		evalFunc: eval,
		PolicyExplore: 3.0,
	}
	return mcts
}

// When a mate in 1 since it is a terminal state win any previously explored branches of the parent are irrelevant, since it will always opt for mate in 1
// So MateAdjust() will return the eval backpropogate to correct this difference and remove all other branches from parent node
// This helps the algorithm find mates exponentially faster than it otherwise would
func MateAdjust(node *MonteCarloNode) float64 {
	for i, child := range node.Parent.Children {
		if node == child {
			// Make the mate the only possible node
			move := node.Parent.Moves[i]
			node.Parent.Children = []*MonteCarloNode{node}
			node.Parent.Moves = []dragontoothmg.Move{move}

			// The value of the parent should be equal to visits since it has won the game
			// So backpropogation will propogate node.Parent.Visits - node.Parent.Value to correct the value up the tree
			// (node.Parent.Visits - node.Parent.Value) + node.Parent.Value = node.Parent.Visits	
		}
	}
	return node.Parent.Visits - node.Parent.Value
}