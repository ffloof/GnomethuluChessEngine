## TODO:

#### Small Things:
    - Remove non queen promotions from movegen code (its so rare it isnt worth wasting computational power on exploring an extreme edge case)
    - Work on enabling move ordering in main branch
        - MVVLVA, but also try promotions and advanced pawn moves early (especially in endgames)
    - Make set position work on 2 plys
    - If a position is extremely imbalanced ie one side is a rook and pawn up or more, skip positional checks and just return material

#### Books
 - End Sygzygygyygygy book
 - Customizeable opening book format

#### Improve time manager
 - Make engine auto play a move if its clearly the best move by a wide margin by n nodes
 - Factor in other players time into time to move
 - Make engine not choose moves only considered last minute (minimum % of nodes?)

#### Engine options
 - Add options for amount of threads to use
 - Add options for multipv lines
 - Add options for pondering on opponents time

#### Machine learning
 - Traditional Eval
 - Traditional Policy
 - Self play policy and eval
 - Time management?
 - Texel Tuning? Contempt?