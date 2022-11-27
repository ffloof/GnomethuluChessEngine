package main

import (
	_"fmt"
	"testing"
	"time"
	"gnomethulu/search"
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"github.com/dylhunn/dragontoothmg"
)

func TestBenchmarks(t *testing.T){
	t.Log("Start Game Positions")
	for _, iters := range []int{10000, 100000, 1000000} {
		searcher := search.NewSearch(policy.UCT, evaluation.Pesto)
		start := time.Now()
		searcher.RunIterations(iters)
		t.Log(int(float64(iters)/time.Since(start).Seconds()), "nps" , iters, "nodes")
	}


	t.Log("Middle Game Positions")
	for _, iters := range []int{10000, 100000, 1000000} {
		searcher := search.NewSearch(policy.UCT, evaluation.Pesto)
		searcher.SetPosition(dragontoothmg.ParseFen("r2q1rk1/pppb1pbp/2n1pnp1/3p4/3PP3/1PN2NP1/PBP2PBP/R2Q1RK1 b - - 0 9"))
		start := time.Now()
		searcher.RunIterations(iters)
		t.Log(int(float64(iters)/time.Since(start).Seconds()), "nps" , iters, "nodes")
	}
}