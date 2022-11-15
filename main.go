package main

import (
	"fmt"	
	"gnomethulu/mcts" 
	//"gnomethulu/uci"
)

func main() {
	//uci.Init(mcts.UCT, mcts.Evaluate)
	
	searcher := mcts.NewSearch(mcts.UCT, mcts.Evaluate)
	searcher.SetPosition("6rk/1p2RQpp/p4p2/5q2/5P1K/6P1/PP5P/8 b - - 0 29")
	//searcher.ApplyStr("e2e4").ApplyStr("e7e5").ApplyStr("f1c4").ApplyStr("b8c6").ApplyStr("d1f3").ApplyStr("b6b7")
	

	fmt.Println(searcher.RunTime(5.0))

	//f3e5, 
	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits, child.Max)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)
}

