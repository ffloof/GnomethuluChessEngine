package search

import (
	"github.com/ffloof/dragontoothmg"
	"time"
	"sort"
	"fmt"
	"math"
)

type MonteCarloTreeSearcher struct {
	startPos dragontoothmg.Board
	Head     *MonteCarloNode
	treeFunc func(*dragontoothmg.Board, dragontoothmg.Move, *[64]int8) float64 //TODO: convert boards to *board
	evalFunc func(*dragontoothmg.Board) float64
	PolicyExplore float64
}

func NewSearch(tree func(*dragontoothmg.Board, dragontoothmg.Move, *[64]int8) float64, eval func(*dragontoothmg.Board) float64) *MonteCarloTreeSearcher {
	board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	freshHead := newNode(nil, &board)

	return &MonteCarloTreeSearcher{
		startPos: board,
		Head:     &freshHead,
		treeFunc: tree,
		evalFunc: eval,
		PolicyExplore: 0.7,
	}
}

func (mcts *MonteCarloTreeSearcher) GetBestMove() dragontoothmg.Move {
	bestIndex := 0
	bestAverage := -1.0
	for i, v := range mcts.Head.Children {
		average := -v.Value/v.Visits
		if average > bestAverage {
			bestIndex = i
			bestAverage = average
		}
	}
	return mcts.Head.Moves[bestIndex]
}


// TODO: get this to work for two plys as well
func (mcts *MonteCarloTreeSearcher) SetPosition(nextState dragontoothmg.Board){
	for i := range mcts.Head.Children {
		move := mcts.Head.Moves[i]
		testBoard := mcts.startPos
		testBoard.Apply(move)
		if nextState == testBoard {
			// Can use information from pondering / previous move analysis
			mcts.startPos = nextState
			mcts.Head = &mcts.Head.Children[i]
			mcts.Head.Parent = nil
			return
		}
	}

	mcts.startPos = nextState
	freshHead := newNode(nil, &mcts.startPos)
	mcts.Head = &freshHead
}

func (mcts *MonteCarloTreeSearcher) PlayingWhite() bool {
	return mcts.startPos.Wtomove
}


func (mcts *MonteCarloTreeSearcher) RunIterations(n int) {
	for i := 0; i < n; i++ {
		mcts.iteration()
	}
}

const MAXNODES = 2000000

func (mcts *MonteCarloTreeSearcher) runTime(seconds float64, stopSignal chan bool) bool {
	start := time.Now()
	for true {
		mcts.RunIterations(10000)
		mcts.PrintInfo()
		elapsed := time.Since(start)
		
		if len(stopSignal) != 0 || MAXNODES < int(mcts.Head.Visits) {
			return true
		}

		if float64(elapsed.Milliseconds()) / 1000 > seconds  {
			break
		}
		
	}
	return false
}

func (mcts *MonteCarloTreeSearcher) TimeManager(bank float64, increment float64, stopSignal chan bool) {
	// Only one move can be played
	if len(mcts.Head.Moves) == 1 {
		fmt.Println("bestmove", mcts.Head.Moves[0].String())
		return
	}

	// Mate in one detection
	for _, move := range mcts.Head.Moves {
		board := mcts.startPos
		board.Apply(move)
		moves := board.GenerateLegalMoves()
		if len(moves) == 0 {
			if board.OurKingInCheck() {
				fmt.Println("bestmove", move.String())
			}
		}
	}

	// Time manager protocol: Run for at least the increment then run for xln(x)/c
	defer func(){
		move := mcts.GetBestMove()
		fmt.Println("bestmove", move.String())
	}()
	

	if increment != 0.0 { 
		if mcts.runTime(increment, stopSignal) { return } 
	}
	if bank != 0.0 { 
		if mcts.runTime(bank*math.Log(bank)/200, stopSignal) { return }
	}
}

func inverseSigmoid(n float64) float64 {
	SigmoidScale := 0.9
	SigmoidCurve := 0.25
	return -100 * math.Log(((2 *SigmoidScale)/(n+SigmoidScale))-1)/SigmoidCurve
}

func (mcts *MonteCarloTreeSearcher) PrintInfo() {
	moveMap := map[dragontoothmg.Move]float64{}
	for i, move := range mcts.Head.Moves {
		moveMap[move] = -mcts.Head.Children[i].Value / mcts.Head.Children[i].Visits
	}

	keys := make([]dragontoothmg.Move, 0, len(moveMap))
	for k := range moveMap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return moveMap[keys[i]] > moveMap[keys[j]]
	})

	for i, key := range keys {
		if i + 1 > 3 {
			break
		}
		move := key.String()
		eval := moveMap[key]
		fmt.Println("info nodes", int(mcts.Head.Visits), "multipv", i+1 ,"score cp", int(inverseSigmoid(eval)), "pv", move)
	}
}