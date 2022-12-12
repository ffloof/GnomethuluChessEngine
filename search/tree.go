package search

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
)

//TODO: compare with a map based implementation
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
			if !node.Expanded {
				node.Moves = board.GenerateLegalMoves()
				node.Children = make([]*MonteCarloNode, len(node.Moves), len(node.Moves))
				node.Expanded = true
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
		
		// Expand nodes based on uct score

		bestChildIndex := 0
		bestScore := -1.0
		parentConstant := mcts.PolicyExplore * math.Log(node.Visits)
		for i, child := range node.Children {
			var score float64

			if child == nil {				
				score = mcts.treeFunc(board, node.Moves[i]) * math.Sqrt(parentConstant/mcts.EmptyVisits)
			} else {
				score = (-child.Value / child.Visits) + (mcts.treeFunc(board, node.Moves[i]) * math.Sqrt(parentConstant/child.Visits))
			}
			if score > bestScore {
				bestScore = score
				bestChildIndex = i
			}
		}

		if node.Children[bestChildIndex] == nil {
			board.Apply(node.Moves[bestChildIndex])

			nextNode := newNode(node, board)
			node.Children[bestChildIndex] = nextNode

			node = nextNode
			evaluation = mcts.evalFunc(board)

			break selectionLoop
		} else {
			board.Apply(node.Moves[bestChildIndex])
			node = node.Children[bestChildIndex]
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

func newNode(parent *MonteCarloNode, board *dragontoothmg.Board) *MonteCarloNode {
	if parent != nil {
		return &MonteCarloNode{
			Parent:   parent,
		}
	} else {
		moves := board.GenerateLegalMoves()
		return &MonteCarloNode{
			Parent: parent,
			Moves: moves,
			Children: make([]*MonteCarloNode, len(moves), len(moves)),
		}
	}
}

type MonteCarloNode struct { //TODO: look into if we can garbage collect some nodes or at least node.Moves
	Parent   *MonteCarloNode
	Children []*MonteCarloNode //TODO: try to mitigate pointer chasing
	Moves    []dragontoothmg.Move
	Value    float64
	Visits   float64
	Expanded bool
}

func NewSearch(tree func(*dragontoothmg.Board, dragontoothmg.Move) float64, eval func(*dragontoothmg.Board) float64) *MonteCarloTreeSearcher {
	board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	mcts := MonteCarloTreeSearcher{
		startPos: board,
		Head:     newNode(nil, &board),
		treeFunc: tree,
		evalFunc: eval,
		PolicyExplore: 0.5,
		EmptyVisits: 0.001,
	}
	return &mcts
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
			node.Parent.Expanded = true

			// The value of the parent should be equal to visits since it has won the game
			// So backpropogation will propogate node.Parent.Visits - node.Parent.Value to correct the value up the tree
			// (node.Parent.Visits - node.Parent.Value) + node.Parent.Value = node.Parent.Visits	
		}
	}
	return node.Parent.Visits - node.Parent.Value
}