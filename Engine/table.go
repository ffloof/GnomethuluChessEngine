package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Entry struct { //TODO: should be 13 bytes write test for struct size
	hash uint64
	BestMove dragontoothmg.Move
	Score int16
	Depth int8
}


type TranspositionTable []Entry

func (table TranspositionTable) Get(board *dragontoothmg.Board) *Entry {
	entry := &table[board.Hash() % uint64(len(table))]
	if entry.hash == board.Hash() {
		return entry
	} else {
		return nil
	}
}

func (table TranspositionTable) Set(board *dragontoothmg.Board, bestMove dragontoothmg.Move, score int16, depth int8) {
	//TODO: is this the best collision behavior
	entry := &table[board.Hash() % uint64(len(table))]
	if entry.hash == 0 || depth >= entry.Depth {
		table[board.Hash() % uint64(len(table))] = Entry{board.Hash(), bestMove, score, depth}
	}
}

func (table TranspositionTable) EmptyPercent() float64 {
	emptyCount := 0
	for i := 0; i < len(table); i++ {
		if table[i].hash == 0 {
			emptyCount += 1
		}
	}
	return float64(emptyCount) / float64(len(table))
}

//TODO: store bound and type of bound for all nodes, this lets us cut more efficiently
//TODO: add two bucket system