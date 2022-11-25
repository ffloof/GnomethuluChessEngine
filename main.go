package main

import (
	"gnomethulu/evaluation"
	"gnomethulu/policy"
	"gnomethulu/uci"
)

func main() {	
	uci.Init(policy.UCT, evaluation.Pesto)
}

