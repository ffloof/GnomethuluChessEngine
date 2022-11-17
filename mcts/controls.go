package mcts

import (
	"github.com/dylhunn/dragontoothmg"
	"time"
	"sort"
	"fmt"
	"math"
)

func inverseSigmoid(n float64) float64 {
	SigmoidScale := 0.9
	SigmoidCurve := 0.25
	return -math.Log(((2 *SigmoidScale)/(n+SigmoidScale))-1)/SigmoidCurve
}



func (mcts *MonteCarloTreeSearcher) SetPosition(nextState dragontoothmg.Board){
	for i, move := range mcts.Head.Moves {
		testBoard := mcts.BaseState
		testBoard.Apply(move)
		if nextState == testBoard {
			// Can use information from pondering / previous move analysis
			mcts.BaseState = nextState
			mcts.Head = mcts.Head.Children[i]
			mcts.Head.Parent = nil
			return
		}
	}

	mcts.BaseState = nextState
	mcts.Head = newNode(nil, mcts.BaseState)
}





func (mcts MonteCarloTreeSearcher) RunIterations(n int) {
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

var hard_breakpoint int = 1000000

func (mcts MonteCarloTreeSearcher) RunInfinite(stop chan bool) {
	for true {
		if len(stop) != 0{
			return
		}
		mcts.RunIterations(10000)

		if int(mcts.Head.Visits) > hard_breakpoint {
			leave := <- stop
			leave = leave
			return
		}

		
		moveMap := map[dragontoothmg.Move]float64{}
		for i, move := range mcts.Head.Moves {
			moveMap[move] = mcts.Head.Children[i].Value / mcts.Head.Children[i].Visits
		}

		keys := make([]dragontoothmg.Move, 0, len(moveMap))
		for k := range moveMap {
			keys = append(keys, k)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return moveMap[keys[i]] > moveMap[keys[j]]
		})

		for i, key := range keys {
			if i + 1 > 3{
				break
			}
			move := key.String()
			eval := moveMap[key]
			fmt.Println("info nodes", mcts.Head.Visits ,"multipv", i+1 ,"score cp", int(inverseSigmoid(eval) * 100), "pv", move)
		}		
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

type MonteCarloTreeSearcher struct {
	BaseState dragontoothmg.Board
	Head      *MonteCarloNode
	treeFunc  func(*MonteCarloNode, *MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64
	evalFunc  func(dragontoothmg.Board) float64
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