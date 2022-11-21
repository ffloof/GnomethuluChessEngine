package main

import (
	"fmt"
	"gnomethulu/engine"
	"github.com/dylhunn/dragontoothmg"
)

func main(){
	startpos := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	var maxDepth int8 = 9
	fmt.Println(engine.NegaMax(&startpos,-9999,9999,maxDepth))
	
	var i int8
	for i=maxDepth;i>=0;i-- {
		fmt.Println("DEPTH",maxDepth - i,":",engine.DepthCount[i])
	}
	fmt.Println(engine.TTable.EmptyPercent())
}

/* 
Without ab pruning
DEPTH 5 : 1                                                                                                                                                                 
DEPTH 4 : 20                                                                                                                                                                 
DEPTH 3 : 400                                                                                                                                                                
DEPTH 2 : 8902                                                                                                                                                               
DEPTH 1 : 197281                                                                                                                                                             
DEPTH 0 : 4865609

With Ab pruning
DEPTH 5 : 1                                                                                                                                                                  
DEPTH 4 : 20                                                                                                                                                                 
DEPTH 3 : 58                                                                                                                                                                 
DEPTH 2 : 844                                                                                                                                                                
DEPTH 1 : 1891                                                                                                                                                               
DEPTH 0 : 35152
*/

/*
Without transposition table
DEPTH 0 : 1                                                                                                                                                                  
DEPTH 1 : 20                                                                                                                                                                 
DEPTH 2 : 39                                                                                                                                                                 
DEPTH 3 : 471                                                                                                                                                                
DEPTH 4 : 1070                                                                                                                                                               
DEPTH 5 : 20528                                                                                                                                                              
DEPTH 6 : 63415                                                                                                                                                              
DEPTH 7 : 1176796

With transposition table
DEPTH 0 : 1
DEPTH 1 : 20                                                                                                                                
DEPTH 2 : 39                                                                                                                                
DEPTH 3 : 476                                                                                                                               
DEPTH 4 : 1087                                                                                                                              
DEPTH 5 : 15504                                                                                                                             
DEPTH 6 : 49579                                                                                                                             
DEPTH 7 : 785615 
*/

/* As full window
DEPTH 0 : 1                                                                                                                                 
DEPTH 1 : 20                                                                                                                                
DEPTH 2 : 41                                                                                                                                
DEPTH 3 : 515                                                                                                                               
DEPTH 4 : 1814                                                                                                                              
DEPTH 5 : 21270                                                                                                                             
DEPTH 6 : 192302                                                                                                                            
DEPTH 7 : 1119110                                                                                                                           
DEPTH 8 : 12721879

As null window
DEPTH 0 : 1                                                                                                                                 
DEPTH 1 : 20                                                                                                                                
DEPTH 2 : 20                                                                                                                                
DEPTH 3 : 445                                                                                                                               
DEPTH 4 : 380                                                                                                                               
DEPTH 5 : 6752                                                                                                                              
DEPTH 6 : 6226                                                                                                                              
DEPTH 7 : 122476                                                                                                                            
DEPTH 8 : 594904
*/