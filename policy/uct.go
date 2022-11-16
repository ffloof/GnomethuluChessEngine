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
	return (child.Value / child.Visits) + (capture_multiplier * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}

func abs(n float64) float64 {
	if n < 0.0 {
		return -n
	}
	return n
}


func MM_UCT(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 0.6
	capture_multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		capture_multiplier = PolicyCapture
	}

	action := child.Value/child.Visits

	//Action Bonus: bonus for paths where parent minmax is close to (child action or child minmax) TODO: figure which best

	correctBonus := parent.Max + child.Max + 1.0

	//Explore bonus: for paths where child action and child minmax vary
	
	exploreBonus := 1+((action + child.Max)/2)

	return (action * correctBonus) + (exploreBonus * capture_multiplier  * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}