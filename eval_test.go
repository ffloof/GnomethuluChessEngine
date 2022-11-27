package main

import (
	"testing"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/evaluation/pesto"
	"gnomethulu/evaluation/custom"
)

func TestEvaluation(t *testing.T){
	startBoard := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	t.Log("EvalBalance", pesto.Pesto(startBoard))
	t.Log("EvalBalance", custom.V1(startBoard))
}