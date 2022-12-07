package search

import (
	"github.com/dylhunn/dragontoothmg"
	"math"
)

func (mcts *MonteCarloTreeSearcher) iteration() {
	evaluation := 0.0

	// 1. Selection
	copiedBoard := mcts.startPos
	board := &copiedBoard
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
		if !node.FullyExpanded {
			for i := range node.Children {
				if node.Children[i] == nil {
					// 2. Expansion and Evaluation
					board.Apply(node.Moves[i])

					nextNode := newNode(node, board)
					node.Children[i] = nextNode

					node = nextNode
					evaluation = mcts.evalFunc(board)
					break selectionLoop
				} else if i == len(node.Children) - 1 {
					node.FullyExpanded = true
				}
			}
		}

		bestChildIndex := 0
		bestScore := -1.0
		parentConstant := mcts.PolicyExplore * math.Log(node.Visits)
		for i, child := range node.Children {
			//score := mcts.treeFunc(parentConstant, child, board, node.Moves[i])
			score := (child.Value / child.Visits) + (mcts.treeFunc(board, node.Moves[i]) * math.Sqrt(parentConstant/child.Visits))
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

func newNode(parent *MonteCarloNode, board *dragontoothmg.Board) *MonteCarloNode {
	moves := board.GenerateLegalMoves()
	children := make([]*MonteCarloNode, len(moves), len(moves))

	return &MonteCarloNode{
		Parent:   parent,
		Children: children,
		Moves:    moves,
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
		PolicyExplore: 2.0,
	}
	return mcts
}