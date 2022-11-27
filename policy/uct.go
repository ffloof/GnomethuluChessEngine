package policy

import (
	"gnomethulu/search"
	"github.com/dylhunn/dragontoothmg"
	"math"
)

var PolicyExplore float64 = 2.5

func UCT(parent, child *search.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := PolicyExplore
	return (child.Value / child.Visits) + math.Sqrt(c*math.Log(parent.Visits)/child.Visits)
}