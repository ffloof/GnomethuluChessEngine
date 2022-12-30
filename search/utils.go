package search

import (
	"fmt"
	"math"
)

func (explore MonteCarloNode) Print() string {
	str := fmt.Sprint(int(explore.Visits))
	for i := range explore.Children {
		child := &explore.Children[i]
		average := -child.Value/child.Visits
		str += "\n" + fmt.Sprint(i, " ", explore.Moves[i].String(), " mean: ", round5(average), " visits: ", int(child.Visits), " minmax: ", round5(-child.MinMax))
	}
	return str
}

func round5(n float64) float64 {
	return math.Round(n * 10000) / 10000
}