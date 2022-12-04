package main

import (
	"gnomethulu/evaluation/custom"
	"gnomethulu/policy"
	"gnomethulu/uci"
	"gnomethulu/evaluation/neural"
)

func main() {	
	neural.Init()
	uci.Init(policy.UCT, custom.V1)
}

