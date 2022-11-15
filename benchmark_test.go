package main

import (
	_"fmt"
	"testing"
	"gnomethulu/mcts"
	"gnomethulu/evaluation"
	"gnomethulu/policy"
)

func TestBenchmarks(t *testing.T){
	for _, time := range []int{1,2,5,10} {
		searcher := mcts.NewSearch(policy.UCT, evaluation.Pesto)
		iter := searcher.RunTime(float64(time))
		t.Log(time, "s    ", iter, "nodes")
	}
}