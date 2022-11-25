package main

import (
	"fmt"	
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"gnomethulu/mcts"
	"gnomethulu/uci"
	"github.com/dylhunn/dragontoothmg"
)

func main() {	
	
	searcher := mcts.NewSearch(policy.UCT, evaluation.Pesto)
	searcher.SetPosition(dragontoothmg.ParseFen("1r1qk2r/p1pb1ppp/3b4/4pnB1/3p4/3P2P1/PPP3BP/R2QK1NR b KQk - 1 14"))

	fmt.Println(searcher.RunTime(5.0))

	explore := searcher.Head
	for i, child := range explore.Children {
		average := child.Value/child.Visits
		fmt.Println(i, explore.Moves[i].String(), child.Visits, "mean", average)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	
	uci.Init(policy.UCT, evaluation.Pesto)
}

