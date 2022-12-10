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
	searcher.SetPosition(dragontoothmg.ParseFen("rnbqkb1r/ppp2ppp/8/3pN1nQ/8/8/PPPP1PPP/RNB1KB1R w KQkq - 0 6"))

	searcher.RunIterations(500000)
	t.Log(searcher.Head.Visits)

	explore := searcher.Head
	for i, child := range explore.Children {
		average := -child.Value/child.Visits
		t.Log(i, explore.Moves[i].String(), child.Visits, "mean", average)
	}
	best := searcher.GetBestMove()
	t.Log(best.String())
}

func TestMCTS2(t *testing.T){
	searcher := search.NewSearch(policy.UCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("4rk2/3QRppp/2p5/p1q5/P1P5/6pP/1P3PP1/4R1K1 b - - 0 29"))

	searcher.RunIterations(500000)
	t.Log(searcher.Head.Visits)

	explore := searcher.Head
	for i, child := range explore.Children {
		average := -child.Value/child.Visits
		t.Log(i, explore.Moves[i].String(), child.Visits, "mean", average)
	}
	best := searcher.GetBestMove()
	t.Log(best.String())
}