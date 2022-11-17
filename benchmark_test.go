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
	t.Log("Start Game Positions")
	for _, time := range []int{1,2,5,10} {
		searcher := mcts.NewSearch(policy.UCT, evaluation.Wrapper)
		t.Log(searcher.BaseState.ToFen())
		iter := searcher.RunTime(float64(time))
		t.Log(time, "s    ", iter, "nodes")
	}


	t.Log("Middle Game Positions")
	for _, time := range []int{1,2,5,10} {
		searcher := mcts.NewSearch(policy.UCT, evaluation.Wrapper)
		searcher.SetPosition(dragontoothmg.ParseFen("r2q1rk1/pppb1pbp/2n1pnp1/3p4/3PP3/1PN2NP1/PBP2PBP/R2Q1RK1 b - - 0 9"))
		t.Log(searcher.BaseState.ToFen())
		iter := searcher.RunTime(float64(time))
		t.Log(time, "s    ", iter, "nodes")
	}
}