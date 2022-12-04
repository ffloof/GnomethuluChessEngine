import chess
import chess.engine
import random
import numpy
import sys

# We want to focus on evaluating relatively quiet boards
# Otherwise it can be very confusing with all the tactics when several pieces are hanging
# We can rely on the engine's search to find tactics in the more complicated positions
# Playing captures more often will lead to simpler positions that tend to be quieter

def capture_moves(board):
	moves = []
	for move in board.legal_moves:
		if board.is_capture(move):
			moves.append(move)
	return moves

def random_board(max_depth=100):
	board = chess.Board()
	depth = random.randrange(0, max_depth)
	for _ in range(depth):
		captures = capture_moves(board)

		if len(captures) == 0 or random.randint(1,3) > 1: #2/3 times just play a random move, 1/3 of the time play a random capturing move
			all_moves = list(board.legal_moves)
			random_move = random.choice(all_moves)
		else:
			random_move = random.choice(captures)

		board.push(random_move)
		if board.is_game_over():
			board.pop()
			return board
	return board

sf = chess.engine.SimpleEngine.popen_uci('./stockfish_15_x64_avx2.exe')
def stockfish(board, depth):
	result = sf.analyse(board, chess.engine.Limit(depth=depth))
	return str(result['score'].white())


data = []

amount = int(sys.argv[1])
for i in range(amount):
	print(i+1, "/", amount)
	rb = random_board()
	evaluation = 0

	evaluation = 0
	outcome = rb.outcome()
	if outcome == None:
		evaluation = stockfish(rb,10)
	elif outcome.winner == None:
		evaluation = "+0"
	elif outcome.winner == chess.WHITE:
		evaluation = "#+0"
	else:
		evaluation = "#-0"

	data.append(rb.fen() + "," + str(evaluation) + "\n")

with open('boards.csv','w') as file:
	file.write("fen,eval\n")
	for line in data:
		file.write(line)

sf.quit()