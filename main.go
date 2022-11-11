package main

import (
	"fmt"
	"gnomethulu/mcts"

	"github.com/dylhunn/dragontoothmg"
)

func main() {
	
	//Starts at bottom goes left to right
	fmt.Println(mcts.Evaluate(dragontoothmg.ParseFen("rnb1kb1r/ppp3pp/5p2/3K4/7q/6P1/PPPPP2P/RNBQ1BNR b kq - 1 8")))
	/*
	searcher := mcts.NewSearch(dragontoothmg.ParseFen("rnb1kb1r/ppp2ppp/8/4K3/7q/8/PPPPP1PP/RNBQ1BNR w kq - 1 7"), mcts.UCT, mcts.Evaluate)

	fmt.Println(searcher.RunTime(5.0))

	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)*/
}

