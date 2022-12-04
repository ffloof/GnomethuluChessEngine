package main

import (
	//"gnomethulu/evaluation/traditional"
	"gnomethulu/policy"
	"gnomethulu/uci"
	"gnomethulu/evaluation/neural"
)

func main() {	
	uci.Init(policy.UCT, neural.Init())
}

