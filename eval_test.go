package main

import (
	"testing"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/evaluation"
)

func TestEvaluation(t *testing.T){
	startBoardEval := evaluation.Pesto(dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
	t.Log("EvalBalance", startBoardEval)
	startBoardEval2 := evaluation.PestoQuiescenceWrapper(dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
	t.Log("EvalBalance 2", startBoardEval2)

}