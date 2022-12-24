package main

import (
	"gnomethulu/evaluation/traditional"
	"gnomethulu/policy"
	"gnomethulu/uci"
	//"gnomethulu/evaluation/neural"
	//"gnomethulu/tournament"
)

func main() {
	//tournament.Run()

	uci.Init(policy.HeurUCT, traditional.CustomV2)
}


