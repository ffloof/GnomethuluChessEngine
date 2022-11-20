package main

import (
	"fmt"
	"gnomethulu/engine"
	"github.com/dylhunn/dragontoothmg"
)

func main(){
	startpos := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	moves := startpos.GenerateLegalMoves()
	for _, move := range moves {
		undo := startpos.Apply(move)
		fmt.Println(move.String(), engine.NegaMax(&startpos, -10000, 10000, 7))
		undo()
	}
}