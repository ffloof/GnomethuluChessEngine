package policy

import (
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

var PolicyExplore float64 = 2.0

func UCT(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := PolicyExplore
	return (child.Value / child.Visits) + (5-(5*(child.Variance/child.Visits)))*math.Sqrt(c*math.Log(parent.Visits)/child.Visits)
}