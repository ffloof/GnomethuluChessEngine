package policy

import (
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

var PolicyExplore float64 = 2.0
var PolicyCapture float64 = 1.5

func UCT(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := PolicyExplore
	multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		multiplier = PolicyCapture
	}
	return (child.Value / child.Visits) + (multiplier * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}