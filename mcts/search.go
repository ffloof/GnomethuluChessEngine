package mcts

import (
	"github.com/dylhunn/dragontoothmg"
	"time"
)

type MonteCarloTreeSearcher struct {
	BaseState dragontoothmg.Board
	Head      *MonteCarloNode
	treeFunc  func(*MonteCarloNode, *MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64
	evalFunc  func(dragontoothmg.Board) float64
}

func NewSearch(tree func(*MonteCarloNode, *MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64, eval func(dragontoothmg.Board) float64) MonteCarloTreeSearcher {
	board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return MonteCarloTreeSearcher{
		BaseState: board,
		Head:      newNode(nil, board),
		treeFunc:  tree,
		evalFunc:  eval,
	}
}

func (mcts *MonteCarloTreeSearcher) SetPosition(fen string){
	mcts.BaseState = dragontoothmg.ParseFen(fen)
	mcts.Head = newNode(nil, mcts.BaseState)
}

func (mcts *MonteCarloTreeSearcher) ApplyMove (nextMove dragontoothmg.Move) *MonteCarloTreeSearcher {
	mcts.BaseState.Apply(nextMove)
	var nextNode *MonteCarloNode
	for i, move := range mcts.Head.Moves {
		if move == nextMove{
			nextNode = mcts.Head.Children[i]
			break
		}
	}

	if nextNode == nil {
		mcts.Head = newNode(nil, mcts.BaseState)
	} else {
		nextNode.Parent = nil
		mcts.Head = nextNode
	}
	return mcts
}

func (mcts *MonteCarloTreeSearcher) ApplyStr (movestr string) *MonteCarloTreeSearcher {
	move, err := dragontoothmg.ParseMove(movestr)
	if err == nil {
		return mcts.ApplyMove(move)
	} else {
		return nil
	}
}

func (mcts MonteCarloTreeSearcher) RunIterations(n int) {
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

func (mcts MonteCarloTreeSearcher) RunTime(seconds float64) int {
	start := time.Now()
	count := 0
	for true {
		mcts.RunIterations(1000)
		count += 1000
		elapsed := time.Since(start)
		if float64(elapsed.Milliseconds()) / 1000 > seconds  {
			break
		}
	}
	return count
}

func (mcts MonteCarloTreeSearcher) GetBestMove() dragontoothmg.Move {
	bestIndex := 0
	bestAverage := -1.0
	for i, v := range mcts.Head.Children {
		if v.Value/v.Visits > bestAverage {
			bestIndex = i
			bestAverage = v.Value / v.Visits
		}
	}
	return mcts.Head.Moves[bestIndex]
}