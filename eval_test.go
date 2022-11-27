package main

import (
	"testing"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/evaluation/v1"
)

func TestEvaluation(t *testing.T){
	startBoardEval := v1.Evaluate(dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
	t.Log("EvalBalance", startBoardEval)
}