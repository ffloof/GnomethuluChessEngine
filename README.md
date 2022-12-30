## TODO:

#### Small Things:
    - Add castling, and pawn moves based on how far advanced to move ordering

#### Books
 - End Sygzygygyygygy book
 - Customizeable opening book format

#### Improve time manager
 - Factor in other players time into time to move
 - Make engine not choose moves only considered last minute (minimum % of nodes?)

#### Engine options
 - Add options for amount of threads to use
 - Add options for multipv lines
 - Add options for pondering on opponents time

#### Machine learning
 - policy and eval
 - Time management?

 #### Hand Crafted Evaluation Aspects
    - Tapered evaluation using phase
    - Contempt/simplification, when significantly ahead materially bonus is applied as phase decreases to encourage simplifying
    - Common endgames, ie insufficient material, or simple king and rook / king and queen
    - Control/Mobility, each square has a bonus associated with controlling it
    - King Safety, based off control of squares adjacent to king, decreases with phase.
    - Pawn structure
        - Passed pawns bonus
        - Isolated pawns bonus
    - Piece square table, static bonuses to encourage certain locations as necessary last to tune
    - Large penalty for if side moving's king is in check