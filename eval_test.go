package main

import (
	"testing"
	"github.com/ffloof/dragontoothmg"
	"gnomethulu/evaluation/traditional"
)

func TestEvaluation(t *testing.T){
	startBoard := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	t.Log("EvalBalance", traditional.Pesto(&startBoard))
	t.Log("EvalBalance", traditional.CustomV1(&startBoard))
}