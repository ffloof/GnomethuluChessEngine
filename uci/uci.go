package uci

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"gnomethulu/search"
	"github.com/dylhunn/dragontoothmg"
)

//TODO: replace all these nasty long func expressions with custom types
func Init(treeFunc func(*dragontoothmg.Board, dragontoothmg.Move) float64, evalFunc func(*dragontoothmg.Board) float64){
	reader := bufio.NewReader(os.Stdin)

	searcher := search.NewSearch(treeFunc, evalFunc)
	stop := make(chan bool, 2)

	f, _ := os.Create("./input_log.txt")
 	w := bufio.NewWriter(f)

	for true {
		line, _ := reader.ReadString('\n')
		w.WriteString(line)
		w.Flush()
		line = strings.TrimSpace(line)
		//TODO: we should make all white space a single space character to simplify edge cases

		command := getFirstWord(line)
		arguments := getStringAfter(line, " ")

		arguments = arguments
		switch (command) {
			case "uci":
				fmt.Println("id name Gnomethulu")
				fmt.Println("id author ffloof")
				fmt.Println("uciok")
			case "isready":
				fmt.Println("readyok")
			case "quit":
				return
			case "ucinewgame":
				searcher = search.NewSearch(treeFunc, evalFunc)
			case "position":
				//first word wont work cause fen contains spaces
				
				fenString := getStringBefore(arguments, "moves")				
				if fenString == "" {
					fenString = "startpos"
				}

				moveStrings := strings.Split(strings.TrimSpace(getStringAfter(arguments, "moves")), " ")
				
				var board dragontoothmg.Board
				if fenString != "startpos" {
					board = dragontoothmg.ParseFen(fenString)
				} else {
					board = dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
				}
				for _, movestr := range moveStrings {
					if movestr != "" {
						move, err := dragontoothmg.ParseMove(movestr)
						if err != nil {
							fmt.Println("info string", err, movestr)
						}
						board.Apply(move)
					}
				}
				searcher.SetPosition(board)
			case "go":
				base := 0.0
				increment := 0.0
				if getFirstWord(arguments) == "infinite" {
					increment = 100000
				} else if getStringAfter(arguments, "movetime") != "" {
					increment = convertFloat(getFirstWord(getStringAfter(arguments, "movetime")))
				} else {
					if searcher.PlayingWhite() {
						base = convertFloat(getFirstWord(getStringAfter(arguments, "wtime")))
						increment = convertFloat(getFirstWord(getStringAfter(arguments, "winc")))
					} else {
						base = convertFloat(getFirstWord(getStringAfter(arguments, "btime")))
						increment = convertFloat(getFirstWord(getStringAfter(arguments, "binc")))
					}
				}

				base /= 1000
				increment /= 1000

				fmt.Println(base, increment)


				stop <- true
				stop = make(chan bool, 2)
				go searcher.TimeManager(base, increment, stop)
				
			case "stop":
				stop <- true
		}
	}
}

func getStringAfter(str string, find string) string {
	start := strings.Index(str, find)
	if start == -1 {
		return ""
	}
	return strings.TrimSpace(str[start + len(find):])
}

func getStringBefore(str string, find string) string {
	end := strings.Index(str, find)
	if end == -1 {
		end = len(str)
	}
	return strings.TrimSpace(str[:end])
}

func getFirstWord(str string) string {
	end := strings.Index(str, " ")
	if end == -1 {
		end = len(str)
	}
	return strings.TrimSpace(str[:end])
}

func convertFloat(str string) float64 {
	str = strings.TrimSpace(str)
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0.0
	} else {
		return float64(number)
	}
}