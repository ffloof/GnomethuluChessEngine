package main

import (
	"testing"
	"github.com/dylhunn/dragontoothmg"
	"gnomethulu/mcts"
)

func TestEvaluation(t *testing.T){
	startBoardEval := mcts.Evaluate(dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
	t.Log("EvalBalance", startBoardEval)
}