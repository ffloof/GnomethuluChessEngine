package main

import (
	"fmt"
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
)

func main(){
	startpos := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	searcher := mcts.NewSearch(startpos)
	searcher.RunIterations(1000000)
	
	explore := searcher.Head
	for i, child := range explore.Children {
		fmt.Println(i, explore.Moves[i].String(), child.Visits, child.Value/child.Visits)
	}
	a := searcher.GetBestMove()
	fmt.Println(a.String())
}

	