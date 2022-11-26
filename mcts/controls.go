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

func (mcts MonteCarloTreeSearcher) PrintSearchIdeas() {
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
		fmt.Println("info nodes", int(mcts.Head.Visits) ,"multipv", i+1 ,"score cp", int(inverseSigmoid(eval) * 100), "pv", move)
	}
}



func (mcts *MonteCarloTreeSearcher) SetPosition(nextState dragontoothmg.Board){
	for i, move := range mcts.Head.Moves {
		testBoard := mcts.startPos
		testBoard.Apply(move)
		if nextState == testBoard && mcts.Head.Children[i] != nil {
			// Can use information from pondering / previous move analysis
			mcts.startPos = nextState
			mcts.Head = mcts.Head.Children[i]
			mcts.Head.Parent = nil
			return
		}
	}

	mcts.startPos = nextState
	mcts.Head = newNode(nil, mcts.startPos)
}


func (mcts MonteCarloTreeSearcher) RunIterations(n int) {
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

func (mcts MonteCarloTreeSearcher) RunTime(seconds float64, stopSignal chan bool) bool {
	start := time.Now()
	for true {
		mcts.RunIterations(10000)
		mcts.PrintSearchIdeas()
		elapsed := time.Since(start)
		
		if len(stopSignal) != 0 {
			return true
		}

		if float64(elapsed.Milliseconds()) / 1000 > seconds  {
			break
		}
		
	}
	return false
}

func (mcts MonteCarloTreeSearcher) TimeManager(bank float64, increment float64, stopSignal chan bool) {
	// Only one move can be played
	if len(mcts.Head.Moves) == 1 {
		fmt.Println("bestmove", mcts.Head.Moves[0].String())
		return
	}

	// Mate in one detection
	for _, move := range mcts.Head.Moves {
		board := mcts.startPos
		moves := board.GenerateLegalMoves()
		if len(moves) == 0 {
			if board.OurKingInCheck() {
				fmt.Println("bestmove", move.String())
			}
		}
	}

	/*
	Time manager setup

	Run for at least the increment then decide what to do next
	Otherwise it will use xln(x)/c
	*/

	defer func(){
		move := mcts.GetBestMove()
		fmt.Println("bestmove", move.String())
	}()
	

	if increment != 0.0 { 
		if mcts.RunTime(increment, stopSignal) { return } 
	}
	if bank != 0.0 { 
		if mcts.RunTime(bank*math.Log(bank)/200, stopSignal) { return }
	}
}



type MonteCarloTreeSearcher struct {
	startPos dragontoothmg.Board
	Head      *MonteCarloNode
	treeFunc  func(*MonteCarloNode, *MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64
	evalFunc  func(dragontoothmg.Board) float64
}

func (mcts MonteCarloTreeSearcher) PlayingWhite() bool {
	return mcts.startPos.Wtomove
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