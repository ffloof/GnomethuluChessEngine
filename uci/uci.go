package uci

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"gnomethulu/mcts"
	"github.com/dylhunn/dragontoothmg"
)

func Init(treeFunc func(*mcts.MonteCarloNode, *mcts.MonteCarloNode, dragontoothmg.Board, dragontoothmg.Move) float64, evalFunc func(dragontoothmg.Board) float64){
	reader := bufio.NewReader(os.Stdin)

	searcher := mcts.NewSearch(treeFunc, evalFunc)

	f, _ := os.Create("./input_log.txt")
 	w := bufio.NewWriter(f)

	for true {
		line, _ := reader.ReadString('\n')
		w.WriteString(line)
		w.Flush()
		line = strings.TrimSpace(line)
		//TODO: we should make all white space a single space character to simplify edge cases

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
				fenString := strings.TrimSpace(GetStringBefore(arguments, "moves"))
				moveStrings := strings.Split(strings.TrimSpace(GetStringAfter(arguments, "moves")), " ")
				if fenString != "startpos" {
					searcher.SetPosition(fenString)
				} else {
					searcher.SetPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
				}
				for _, move := range moveStrings {
					searcher.ApplyStr(move)
				}
			case "go":
				timethinker := GetStringAfter(arguments,"movetime")
				if timethinker != "" {
					i, _ := strconv.Atoi(strings.TrimSpace(timethinker))
					searcher.RunTime(float64(i)/1000)
					move := searcher.GetBestMove()
					fmt.Println("bestmove", move.String())
				}

				//TODO: use channels to let ai respond while thinking
				//TODO: constantly transmit data
			case "stop":
				move := searcher.GetBestMove()
				fmt.Println("bestmove", move.String())
				fmt.Println("info string", "stop", arguments)
		}
		
	}
}

func GetStringBefore(str string, end string) string {
	e := strings.Index(str, end)
	if e == -1 {
		return ""
	}
	return str[:e]
}

func GetStringAfter(str string, start string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	return str[s+len(start):]
}