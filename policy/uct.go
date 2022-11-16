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
	multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		multiplier = PolicyCapture
	}
	return (child.Value / child.Visits) + (multiplier * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}

func abs(n float64) float64 {
	if n < 0.0 {
		return -n
	}
	return n
}

func max1(n float64) float64 {
	if n > 1 {
		return 1
	}
	return n
}

func MM_UCT1(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 0.5
	multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		multiplier = PolicyCapture
	}

	winrate := (child.Value/child.Visits)
	difference := 0.0
	if winrate > -child.Max {
		difference = abs(child.Max + winrate)
	}

	if difference != 0.0 {
		if child.Visits > 100 {
			return winrate + (difference) + (multiplier *(math.Sqrt(c*math.Log(parent.Visits)/child.Visits)))
		} else {
			if parent.Max == -child.Max {
				return 100
			}
		}
	}
	return winrate + (multiplier *(math.Sqrt(c*math.Log(parent.Visits)/child.Visits)))
}


func MM_UCT2(parent, child *mcts.MonteCarloNode, parentBoard dragontoothmg.Board, move dragontoothmg.Move) float64 {
	c := 0.5
	multiplier := 1.0
	if dragontoothmg.IsCapture(move, &parentBoard) {
		multiplier = PolicyCapture
	}

	//Correct Bonus: bonus for paths where parent minmax is close to (child action or child minmax)
	//Surprise Bonus: bonus for paths where child action and child minmax vary

	multiplier_2 := 2.0+abs((child.Value/child.Visits) + child.Max)
	

	return (child.Value / child.Visits) + (multiplier_2 * multiplier * math.Sqrt(c*math.Log(parent.Visits)/child.Visits))
}