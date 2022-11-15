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
	stop := make(chan bool)

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
				fmt.Println("option name policyexplore type spin default 20 min 1 max 200")
				fmt.Println("option name policycapture type spin default 150 min 0 max 200")
				fmt.Println("option name sigmoidscale type spin default 90 min 1 max 200")
				fmt.Println("option name sigmoidcurve type spin default 30 min 1 max 200")
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
				if fenString == "" {
					fenString = "startpos"
				}

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
				if strings.Contains(arguments, "infinite") {
					stop = make(chan bool)
					go searcher.RunInfinite(stop)
				} else {
					timethinker := GetStringAfter(arguments,"movetime")
					if timethinker != "" {
						moves := searcher.BaseState.GenerateLegalMoves()
						if len(moves) == 0 {
							fmt.Println("bestmove", moves[0].String())
						} else {
							i, _ := strconv.Atoi(strings.TrimSpace(timethinker))
							searcher.RunTime(float64(i)/1000)
							move := searcher.GetBestMove()
							fmt.Println("bestmove", move.String())
						}
					} else {
						// TODO:  Use time controls
						searcher.RunTime(5.0)
						move := searcher.GetBestMove()
						fmt.Println("bestmove", move.String())
					}
				}
				//TODO: use channels to let ai respond while thinking
				//TODO: constantly transmit data
			case "stop":
				stop <- true
				move := searcher.GetBestMove()
				fmt.Println("bestmove", move.String())
			case "setoption":
				peStr := GetStringAfter(arguments, "policyexplore")
				if peStr != "" {
					temp, err := strconv.Atoi(peStr)
					if err == nil {
					mcts.PolicyExplore = float64(temp)/100
					}
				}
				pcStr := GetStringAfter(arguments, "policycapture")
				if pcStr != "" {
					temp, err := strconv.Atoi(pcStr)
					if err == nil {
						mcts.PolicyCapture = float64(temp)/100
					}
				}
				ssStr := GetStringAfter(arguments, "sigmoidscale")
				if ssStr != "" {
					temp, err := strconv.Atoi(ssStr)
					if err == nil {
						mcts.SigmoidScale = float64(temp)/100
					}
				}
				scStr := GetStringAfter(arguments, "sigmoidcurve")
				if scStr != "" {
					temp, err := strconv.Atoi(scStr)
					if err == nil {
						mcts.SigmoidCurve = float64(temp)/100
					}
				}
				
				
				
		}
	}
}

func GetStringBefore(str string, end string) string {
	e := strings.Index(str, end)
	if e == -1 {
		return ""
	}
	return strings.TrimSpace(str[:e])
}

func GetStringAfter(str string, start string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	return strings.TrimSpace(str[s+len(start):])
}