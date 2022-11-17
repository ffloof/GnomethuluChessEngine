package policy

import (
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

var PolicyExplore float64 = 1.0
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