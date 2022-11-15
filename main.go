package main

import (
	//"fmt"	
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"gnomethulu/uci"
)

func main() {
	uci.Init(policy.MM_UCT, evaluation.Pesto)
	
	/*
	searcher := mcts.NewSearch(mcts.UCT, mcts.Evaluate)
	//searcher.SetPosition("6rk/1p2RQpp/p4p2/5q2/5P1K/6P1/PP5P/8 b - - 0 29")
	//searcher.ApplyStr("e2e4").ApplyStr("e7e5").ApplyStr("g1f3").ApplyStr("b6b7")
	

	fmt.Println(searcher.RunTime(5.0))

	explore := searcher.Head.Children[17].Children[25].Children[18].Children[31]
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits, -child.Max)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)*/
}

