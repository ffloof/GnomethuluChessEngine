package main

import (
	"testing"
	"gnomethulu/policy"
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/evaluation"
)

func TestMCTS(t *testing.T){
	searcher := mcts.NewSearch(policy.UCT, evaluation.Pesto)
	searcher.SetPosition(dragontoothmg.ParseFen("r1bqkb1r/1pp2ppp/p1p2n2/4N3/4P3/2N5/PPPP1PPP/R1BQK2R b KQkq - 0 6"))

	searcher.RunTime(5.0)
	t.Log(searcher.Head.Visits)

	explore := searcher.Head
	for i, child := range explore.Children {
		average := child.Value/child.Visits
		t.Log(i, explore.Moves[i].String(), child.Visits, "mean", average)
	}
	best := searcher.GetBestMove()
	t.Log(best.String())
}


