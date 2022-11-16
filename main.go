package main

import (
	"fmt"	
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"gnomethulu/mcts"
	//"gnomethulu/uci"
)

func main() {
	//uci.Init(policy.UCT, evaluation.Pesto)
	
	
	searcher := mcts.NewSearch(policy.MM_UCT, evaluation.Pesto)
	searcher.SetPosition("2rqkb1r/ppp1pppp/8/3p1b2/1n1Pn2B/2N1PN2/PPP2PPP/2RQKB1R w Kk - 1 8")
	//searcher.ApplyStr("e2e4").ApplyStr("e7e5").ApplyStr("g1f3").ApplyStr("b6b7")
	

	fmt.Println(searcher.RunTime(5.0))

	//a2a3
	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits, -child.Max)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)
}

