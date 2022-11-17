package main

import (
	_"fmt"
	"testing"
	"gnomethulu/mcts"
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"github.com/dylhunn/dragontoothmg"
)

func TestBenchmarks(t *testing.T){
	for _, time := range []int{1,2,5,10} {
		searcher := mcts.NewSearch(policy.UCT, evaluation.Wrapper)
		searcher.SetPosition(dragontoothmg.ParseFen("r2qkb1r/ppp1pppp/2n2n2/3p2B1/3P2b1/2N2N2/PPP1PPPP/R2QKB1R w KQkq - 4 5"))
		t.Log(searcher.BaseState.ToFen())
		iter := searcher.RunTime(float64(time))
		t.Log(time, "s    ", iter, "nodes")
	}
}