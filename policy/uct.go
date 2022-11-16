package policy

import (
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

var PolicyExplore float64 = 0.5
var PolicyCapture float64 = 1.5

func UCT(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := PolicyExplore
	capture_multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		capture_multiplier = PolicyCapture
	}
	//TODO: see if moving capture_multiplier to first term gives any improvement
	return (child.Value / child.Visits) + (capture_multiplier * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}

func abs(n float64) float64 {
	if n < 0.0 {
		return -n
	}
	return n
}


func MM_UCT(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 2.0
	capture_multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		capture_multiplier = PolicyCapture
	}

	action := child.Value/child.Visits


	// TODO: fix the names because neither name is accurate or true
	correctBonus := 2.0 - (parent.Max + child.Max)
	exploreBonus := 1.0 + abs((action + child.Max)/2)

	return (action * correctBonus * exploreBonus) + (capture_multiplier  * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}