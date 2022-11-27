package main

import (
	"gnomethulu/evaluation/v1"
	"gnomethulu/policy"
	"gnomethulu/uci"
)

func main() {	
	uci.Init(policy.UCT, v1.Evaluate)
}

