package main

import (
	"gnomethulu/engine"
	"github.com/dylhunn/dragontoothmg"
)

//TODO: consider internal iterative deepening
func main(){
	startpos := dragontoothmg.ParseFen("r4rk1/ppp2ppp/2nb3q/3N4/3P2b1/3B1N2/PPP2PPP/R2Q1RK1 w - - 1 11")
	//searcher := engine.NewSearch()
	//fmt.Println(searcher.NegaMax(&startpos, -4, -3, 4))

	engine.Cstar(startpos)
	//engine.Base(startpos)
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

/* Full Quiescence Search
DEPTH 0 : 1                                                                                                                                           
DEPTH 1 : 20                                                                                                                                          
DEPTH 2 : 39                                                                                                                                          
DEPTH 3 : 464                                                                                                                                         
DEPTH 4 : 852                                                                                                                                         
DEPTH 5 : 7776                                                                                                                                        
DEPTH 6 : 21659                                                                                                                                       
DEPTH 7 : 482646                                                                                                                                      
DEPTH 8 : 1454502                                                                                                                                     
DEPTH 9 : 2034476                                                                                                                                     
DEPTH 10 : 670131                                                                                                                                     
DEPTH 11 : 805531                                                                                                                                     
DEPTH 12 : 539496                                                                                                                                     
DEPTH 13 : 482121                                                                                                                                     
DEPTH 14 : 274553                                                                                                                                     
DEPTH 15 : 230777                                                                                                                                     
DEPTH 16 : 136835                                                                                                                                     
DEPTH 17 : 160970                                                                                                                                     
DEPTH 18 : 86525                                                                                                                                      
DEPTH 19 : 96915                                                                                                                                      
DEPTH 20 : 51527                                                                                                                                      
DEPTH 21 : 64587                                                                                                                                      
DEPTH 22 : 32052                                                                                                                                      
DEPTH 23 : 36761                                                                                                                                      
DEPTH 24 : 17041                                                                                                                                      
DEPTH 25 : 17181                                                                                                                                      
DEPTH 26 : 7454                                                                                                                                       
DEPTH 27 : 7782                                                                                                                                       
DEPTH 28 : 3507                                                                                                                                       
DEPTH 29 : 1900                                                                                                                                       
DEPTH 30 : 736                                                                                                                                        
DEPTH 31 : 127                                                                                                                                        
DEPTH 32 : 8                                                                                                                                          
DEPTH 33 : 39                                                                                                                                         
DEPTH 34 : 2                                                                                                                                          
DEPTH 35 : 9

Quiescence Search With Captures Only
DEPTH 0 : 1                                                                                                                                           
DEPTH 1 : 20                                                                                                                                          
DEPTH 2 : 39                                                                                                                                          
DEPTH 3 : 464                                                                                                                                         
DEPTH 4 : 852                                                                                                                                         
DEPTH 5 : 7776                                                                                                                                        
DEPTH 6 : 21656                                                                                                                                       
DEPTH 7 : 482467                                                                                                                                      
DEPTH 8 : 102037                                                                                                                                      
DEPTH 9 : 55315                                                                                                                                       
DEPTH 10 : 42895                                                                                                                                      
DEPTH 11 : 39017                                                                                                                                      
DEPTH 12 : 34761                                                                                                                                      
DEPTH 13 : 23922                                                                                                                                      
DEPTH 14 : 17110                                                                                                                                      
DEPTH 15 : 13269                                                                                                                                      
DEPTH 16 : 9758                                                                                                                                       
DEPTH 17 : 8097                                                                                                                                       
DEPTH 18 : 5553                                                                                                                                       
DEPTH 19 : 4355                                                                                                                                       
DEPTH 20 : 3142                                                                                                                                       
DEPTH 21 : 2563                                                                                                                                       
DEPTH 22 : 1786                                                                                                                                       
DEPTH 23 : 1281                                                                                                                                       
DEPTH 24 : 897                                                                                                                                        
DEPTH 25 : 670                                                                                                                                        
DEPTH 26 : 436                                                                                                                                        
DEPTH 27 : 259                                                                                                                                        
DEPTH 28 : 185                                                                                                                                        
DEPTH 29 : 79                                                                                                                                         
DEPTH 30 : 15                                                                                                                                         
DEPTH 31 : 2                                                                                                                                          
DEPTH 32 : 3                                                                                                                                          
DEPTH 33 : 1                                                                                                                                          
DEPTH 34 : 1
*/

/* Before move ordering middlegame position r2qkb1r/ppp1n1pp/2n2p2/1B1ppb2/3P4/1P2PN2/PBP2PPP/RN1Q1RK1 b kq - 0 7
DEPTH 0 : 1                                                                                                                                                                  
DEPTH 1 : 30                                                                                                                                                                 
DEPTH 2 : 169                                                                                                                                                                
DEPTH 3 : 2423                                                                                                                                                               
DEPTH 4 : 1595                                                                                                                                                               
DEPTH 5 : 1993                                                                                                                                                               
DEPTH 6 : 3329                                                                                                                                                               
DEPTH 7 : 3945                                                                                                                                                               
DEPTH 8 : 8464                                                                                                                                                               
DEPTH 9 : 8066                                                                                                                                                               
DEPTH 10 : 16602                                                                                                                                                             
DEPTH 11 : 10986                                                                                                                                                             
DEPTH 12 : 19312                                                                                                                                                             
DEPTH 13 : 11225                                                                                                                                                             
DEPTH 14 : 16887                                                                                                                                                             
DEPTH 15 : 10837                                                                                                                                                             
DEPTH 16 : 14041                                                                                                                                                             
DEPTH 17 : 9173                                                                                                                                                              
DEPTH 18 : 10272                                                                                                                                                             
DEPTH 19 : 6386                                                                                                                                                              
DEPTH 20 : 5744                                                                                                                                                              
DEPTH 21 : 3359                                                                                                                                                              
DEPTH 22 : 2690                                                                                                                                                              
DEPTH 23 : 1517                                                                                                                                                              
DEPTH 24 : 1025                                                                                                                                                              
DEPTH 25 : 462                                                                                                                                                               
DEPTH 26 : 235                                                                                                                                                               
DEPTH 27 : 55                                                                                                                                                                
DEPTH 28 : 21                                                                                                                                                                
DEPTH 29 : 1


After promotes, mmv_lva, other move ordering                                                                                                                                                                        
DEPTH 0 : 1                                                                                                                                                                  
DEPTH 1 : 30                                                                                                                                                                 
DEPTH 2 : 82                                                                                                                                                                 
DEPTH 3 : 479                                                                                                                                                                
DEPTH 4 : 278                                                                                                                                                                
DEPTH 5 : 192                                                                                                                                                                
DEPTH 6 : 272                                                                                                                                                                
DEPTH 7 : 220                                                                                                                                                                
DEPTH 8 : 172                                                                                                                                                                
DEPTH 9 : 138                                                                                                                                                                
DEPTH 10 : 123                                                                                                                                                               
DEPTH 11 : 113                                                                                                                                                               
DEPTH 12 : 45                                                                                                                                                                
DEPTH 13 : 14                                                                                                                                                                
DEPTH 14 : 3                                                                                                                                                                 
DEPTH 15 : 1                                                                                                                                                                 
*/