package mcts

import (
	"github.com/dylhunn/dragontoothmg"
)

func (mcts MonteCarloTreeSearcher) iteration() {
	// 1. Selection
	board := mcts.BaseState
	node := mcts.Head


selectionLoop:
	for true {
		if len(node.Moves) == 0 {
			//TODO: possibly node=node.Parent goes here
			break selectionLoop
		}
		// If any null node exists expand it otherwise choose the one with best uct score

		for i := range node.Children {
			if node.Children[i] == nil {
				// 2. Expansion and Evaluation
				board.Apply(node.Moves[i])

				nextNode := mcts.newNode(node, board)
				node.Children[i] = nextNode

				node = nextNode
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
	evaluation := -node.Max
	max := node.Max


	for node != nil {
		for _, child := range node.Children {
			if child != nil {
				if -child.Max > max {
					max = -child.Max
				}
			}
		}

		node.Max = max

		node.Visits++
		node.Value += evaluation
		evaluation = -evaluation
		max = -max
		node = node.Parent
	}
}

//TODO: make this part of mcts possibly so evalFunc doesnt need to be passed
func (mcts MonteCarloTreeSearcher) newNode(parent *MonteCarloNode, board dragontoothmg.Board) *MonteCarloNode {
	moves := board.GenerateLegalMoves()

	children := make([]*MonteCarloNode, len(moves), len(moves))
	eval := 0.0
	if len(moves) == 0 {
		if board.OurKingInCheck() {
			eval = -1.0 //Got checkmated
		}
		//Stalemate
	} else {
		eval = mcts.evalFunc(board) //Non terminal
	}

	return &MonteCarloNode{
		Parent:   parent,
		Children: children,
		Moves:    moves,
		Value:    eval,
		Visits:   1.0,
		Max: eval,
	}
}

type MonteCarloNode struct {
	Parent   *MonteCarloNode
	Children []*MonteCarloNode //TODO: consider making this just a list of nodes and run benchmarks
	Moves    []dragontoothmg.Move
	Value    float64 
	Visits   float64
	Max float64
}
