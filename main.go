package main

import (
	"fmt"
	"gnomethulu/engine"
	"github.com/dylhunn/dragontoothmg"
)

func main(){
	s := engine.Searcher {
		make(engine.TranspositionTable, 1024*1024, 1024*1024),
	}

	startpos := dragontoothmg.ParseFen("2rr2k1/pb3p1p/4p1p1/nP2B3/2P2q2/P4N1P/2QnBPP1/3R1RK1 b - - 1 21")
	moves := startpos.GenerateLegalMoves()
	for _, move := range moves {
		undo := startpos.Apply(move)
		fmt.Println(move.String(), -s.NegaMax(&startpos, -10000, 10000, 6))
		undo()
	}
}