package main

import (
	"testing"
	"gnomethulu/policy"
	"gnomethulu/search"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/evaluation/traditional"
)

func TestMCTS(t *testing.T){
	searcher := search.NewSearch(policy.UCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("8/8/2q5/8/8/2k5/8/1K6 b - - 1 1"))

	searcher.RunIterations(500000)
	t.Log(searcher.Head.Visits)

	explore := searcher.Head
	PrettyPrintMoves(explore,t)

	best := searcher.GetBestMove()
	t.Log(best.String())
}

func TestMCTS2(t *testing.T){
	searcher := search.NewSearch(policy.UCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("4rk2/3QRppp/2p5/p1q5/P1P5/6pP/1P3PP1/4R1K1 b - - 0 29"))

	searcher.RunIterations(200000)
	t.Log(searcher.Head.Visits)

	//c5f2, g1h1, f2e1, e7e1
	explore := searcher.Head
	PrettyPrintMoves(explore,t)

	best := searcher.GetBestMove()
	t.Log(best.String())
}

func PrettyPrintMoves(explore *search.MonteCarloNode, t *testing.T){
	for i, child := range explore.Children {
		if child == nil {
			t.Log(i, explore.Moves[i].String(), 0)
		} else {
			average := -child.Value/child.Visits
			t.Log(i, explore.Moves[i].String(), child.Visits, "mean", average)
		}
	}
}