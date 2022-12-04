## TODO:

#### Books
 - End Sygzygygyygygy book
 - Customizeable opening book format
 - Create mate search system, garbage collect nodes

#### Improve time manager
 - Make engine auto play a move if its clearly the best move by a wide margin by n nodes
 - Factor in other players time into time to move
 - Play with reaching a node target
 - Make engine not choose moves only considered last minute (minimum % of recent nodes?)

#### Engine options
 - Add options for amount of threads to use
 - Add options for multipv lines
 - Add options for pondering on opponents time

#### Machine learning
 - Traditional Eval
 - Traditional Policy
 - Self play policy and eval
 - Time management?

#### Other
 - Benchmark comparison with direct storage of nodes as opposed to pointers


## NN structure
    - Train it on quiet or semi quiet positions, completely random positions have too many tactics that are hard to capture in a NN
    - Layers :
        empty? experiment with giving it an empty as a value
        pawn   
        knight
        bishop (queen represented as bishop and rook)
        rook
        knight
        moving color