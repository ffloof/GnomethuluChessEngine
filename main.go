package main

import (
	/*
	"fmt"
	
	"github.com/dylhunn/dragontoothmg"
	*/
	"gnomethulu/mcts"
	"gnomethulu/uci"
)

func main() {
	//fmt.Println(mcts.Evaluate(dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")))
	uci.Init(mcts.UCT, mcts.Evaluate)
	//Starts at bottom goes left to right
	
	/*
	searcher := mcts.NewSearch(mcts.UCT, mcts.Evaluate)
	searcher.ApplyStr("d2d4").ApplyStr("d4d5").ApplyStr("c2c4")

	fmt.Println(searcher.RunTime(5.0))

	//f3e5, 
	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits)
	}
	best := searcher.GetBestMove()
	fmt.Println(best.String())
	searcher.ApplyMove(best)
	*/
}

