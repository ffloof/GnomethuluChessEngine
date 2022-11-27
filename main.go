package main

import (
	"gnomethulu/evaluation/custom"
	"gnomethulu/policy"
	"gnomethulu/uci"
)

func main() {	
	uci.Init(policy.UCT, custom.V1)
}

