package main

import (
	"fmt"
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
)

func main(){
	startpos := dragontoothmg.ParseFen("rnbqkb1r/ppp1pppp/8/8/2BPn3/8/PP3PPP/RNBQK1NR w KQkq - 0 5")
	fmt.Println(mcts.Evaluate(startpos))

	searcher := mcts.NewSearch(startpos)
	searcher.RunIterations(100000)
	test := searcher.GetBestMove()
	fmt.Println(test.String())

}