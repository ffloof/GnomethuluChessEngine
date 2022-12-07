package policy

import (
	"math"
	"gnomethulu/search"
	"github.com/dylhunn/dragontoothmg"
)


func UCT(parent float64, child *search.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	return (child.Value / child.Visits) + math.Sqrt(parent/child.Visits)
}