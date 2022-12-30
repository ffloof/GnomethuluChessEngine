package main

import (
	"testing"
	"gnomethulu/policy"
	"gnomethulu/search"
	"github.com/ffloof/dragontoothmg"
	"gnomethulu/evaluation/traditional"
)

func TestMCTS(t *testing.T){
	searcher := search.NewSearch(policy.HeurUCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("r1bqkb1r/ppp1pppp/2n2n2/3p4/3P1B2/2N5/PPP1PPPP/R2QKBNR w KQkq - 4 4"))

	searcher.RunIterations(500000)

	explore := searcher.Head
	t.Log(explore.Print())

	best := searcher.GetBestMove()
	t.Log(best.String())
}

func TestMCTS2(t *testing.T){
	searcher := search.NewSearch(policy.HeurUCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("4rk2/3QRppp/2p5/p1q5/P1P5/6pP/1P3PP1/4R1K1 b - - 0 29"))

	searcher.RunIterations(100000)

	//c5f2, g1h1, f2e1, e7e1
	explore := searcher.Head
	t.Log(explore.Print())

	best := searcher.GetBestMove()
	t.Log(best.String())
}



func TestMCTS3(t *testing.T){
	searcher := search.NewSearch(policy.HeurUCT, traditional.CustomV1)
	searcher.SetPosition(dragontoothmg.ParseFen("3r2k1/1p1n1pp1/ppr4p/4P3/1PPQ1P2/P1R3Bq/8/3R2K1 w - - 5 34"))

	searcher.RunIterations(1000000)

	explore := searcher.Head
	t.Log(explore.Print())

	best := searcher.GetBestMove()
	t.Log(best.String())
}