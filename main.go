package main

import (
	//"fmt"	
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	//"gnomethulu/mcts"
	"gnomethulu/uci"
)

func main() {
	uci.Init(policy.MM_UCT1, evaluation.Pesto)
	
	/*
	searcher := mcts.NewSearch(policy.MM_UCT, evaluation.Pesto)
	searcher.SetPosition("rnbqkbnr/pppppppp/8/8/8/2N5/PPPPPPPP/R1BQKBNR b KQkq - 1 1")
	//searcher.ApplyStr("e2e4").ApplyStr("e7e5").ApplyStr("g1f3").ApplyStr("b6b7")
	

	fmt.Println(searcher.RunTime(5.0))

	//b7b5, c3b5, a7a6
	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits, -child.Max)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)*/
}

