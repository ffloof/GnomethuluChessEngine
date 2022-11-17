package main

import (
	//"fmt"	
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	//"gnomethulu/mcts"
	"gnomethulu/uci"
	//"github.com/dylhunn/dragontoothmg"
)

func main() {
	uci.Init(policy.UCT, evaluation.PestoQuiescenceWrapper)
	/*
	
	
	
	searcher := mcts.NewSearch(policy.MM_UCT, evaluation.Wrapper)
	searcher.SetPosition("2rq1rk1/2pbbpp1/p1nppn1p/1p6/3PP2B/PBN2N1P/1PP2PP1/2RQ1RK1 b - - 1 12")
	//searcher.ApplyStr("e2e4").ApplyStr("e7e5").ApplyStr("g1f3").ApplyStr("b6b7")
	

	fmt.Println(searcher.RunTime(5.0))

	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits, -child.Max)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)*/
}

