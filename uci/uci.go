package uci

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
)

func Init(treeFunc func(*mcts.MonteCarloNode, *mcts.MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64, evalFunc func(dragontoothmg.Board) float64){
	reader := bufio.NewReader(os.Stdin)

	searcher := mcts.NewSearch(treeFunc, evalFunc)

	for true {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		command := ""
		arguments := ""
		i := strings.Index(line, " ")
		if i == -1 {
			command = line
		} else {
			command = line[0:i]
			arguments = line[i+1:]
		}

		arguments = arguments
		switch (command) {
			case "uci":
				fmt.Println("id name Gnomethulu")
				fmt.Println("id author ffloof")
				fmt.Println("option name Worthless type spin default 1 min 1 max 128")
				//TODO: add options here
				fmt.Println("uciok")
			case "isready":
				fmt.Println("readyok")
			case "quit":
				return
			case "ucinewgame":
				searcher = mcts.NewSearch(treeFunc, evalFunc)
				searcher = searcher
			case "position":

			case "go":
		}
		//fmt.Println(line)
	}
}

/* c2s
uci
debug

*/