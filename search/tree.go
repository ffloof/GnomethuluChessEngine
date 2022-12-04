package search

import (
	"github.com/dylhunn/dragontoothmg"
)

func (mcts *MonteCarloTreeSearcher) iteration() {
	evaluation := 0.0

	// 1. Selection
	board := mcts.startPos
	node := mcts.Head

selectionLoop:
	for true {
		if len(node.Moves) == 0 {
			if board.OurKingInCheck() {
				evaluation = -1.0
			} else {
				evaluation = 0.0
			}
			break selectionLoop
		}
		// If any null node exists expand it otherwise choose the one with best uct score

		for i := range node.Children {
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
		for i, v := range node.Children {
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
		evaluation = -evaluation
		node.Visits++
		node.Value += evaluation
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
	}
}

type MonteCarloNode struct {
	Parent   *MonteCarloNode
	Children []*MonteCarloNode
	Moves    []dragontoothmg.Move
	Value    float64
	Visits   float64
}

func NewSearch(tree func(*MonteCarloNode, *MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64, eval func(dragontoothmg.Board) float64) MonteCarloTreeSearcher {
	board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	mcts := MonteCarloTreeSearcher{
		startPos: board,
		Head:     newNode(nil, board),
		treeFunc: tree,
		evalFunc: eval,
	}
	return mcts
}