package engine

import (
	"fmt"
)

func (search *Searcher) PrintDepths(maxDepth int8){
	i := maxDepth
	for true {
		_, contains := search.DepthCount[i]
		if !contains { break }
		fmt.Println("DEPTH",maxDepth - i,":",search.DepthCount[i])
		i--
	}
	fmt.Println(search.Table.EmptyPercent())
}