package main

import (
	_"fmt"
	"testing"
	"gnomethulu/mcts"
)

func TestBenchmarks(t *testing.T){
	for _, time := range []int{1,2,5,10} {
		searcher := mcts.NewSearch(mcts.UCT, mcts.Evaluate)
		iter := searcher.RunTime(float64(time))
		t.Log(time, "s    ", iter, "nodes")
	}
}