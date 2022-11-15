package main

import (
	"fmt"	
	"gnomethulu/mcts" 
	//"gnomethulu/uci"
)

func main() {
	//uci.Init(mcts.UCT, mcts.Evaluate)
	
	searcher := mcts.NewSearch(mcts.UCT, mcts.Evaluate)
	searcher.SetPosition("rnbqkbnr/1ppp1ppp/p7/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 3")
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

