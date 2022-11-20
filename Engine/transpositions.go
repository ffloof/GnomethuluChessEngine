package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Entry struct {
	hash uint64
	score int16
	best dragontoothmg.Move
	depth int8
	//TODO: prolly add age or something
}

type TranspositionTable struct {
	entries []Entry
}

func NewTranspositionTable(size int) TranspositionTable {
	return TranspositionTable{
		entries : make([]Entry, size, size),
	}
}

//TODO: make sure to add checks and prioritize things
func (table *TranspositionTable) Get(board *dragontoothmg.Board, depth int8) *Entry {
	entry := &table.entries[board.Hash() % uint64(len(table.entries))]
	if board.Hash() == entry.hash  {
		return entry
	} else {
		return nil
	}
}

func (table *TranspositionTable) Set(board *dragontoothmg.Board, depth int8, score int16, best dragontoothmg.Move) {
	table.entries[board.Hash() % uint64(len(table.entries))] = Entry{board.Hash(), score, best, depth}
} 
