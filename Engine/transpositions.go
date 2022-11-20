package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Entry struct {
	hash uint64
	Score int16
	BestMove dragontoothmg.Move
	Depth int8
	//TODO: prolly add age or something
}

//TODO: make this a struct with a rw mutex
type TranspositionTable []Entry

//TODO: make sure to add checks and prioritize things
func (table TranspositionTable) Get(board *dragontoothmg.Board) *Entry {
	entry := &table[board.Hash() % uint64(len(table))]
	if board.Hash() == entry.hash  {
		return entry
	} else {
		return nil
	}
}

func (table TranspositionTable) Set(board *dragontoothmg.Board, depth int8, score int16, best dragontoothmg.Move) {
	table[board.Hash() % uint64(len(table))] = Entry{board.Hash(), score, best, depth}
} 
