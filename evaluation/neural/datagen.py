import chess
import chess.engine
import random
import numpy

def random_board(max_depth=100):
	board = chess.Board()
	depth = random.randrange(0, max_depth)
	for _ in range(depth):
		all_moves = list(board.legal_moves)
		random_move = random.choice(all_moves)
		board.push(random_move)
		if board.is_game_over():
			break
	return board

def stockfish(board, depth):
	with chess.engine.SimpleEngine.popen_uci('./stockfish_15_x64_avx2.exe') as sf:
		result = sf.analyse(board, chess.engine.Limit(depth=depth))
		score = result['score'].white().score()
		return score

rb = random_board()

def convert(board):
	inputlayer = numpy.zeros((6, 8, 8), dtype=numpy.int8)

	for i in range(0,64):
		piece = rb.piece_at(i)
		if piece != None:
			y = i//8
			x = i%8
			inputlayer[0][y][x] = piece.color == board.turn
			inputlayer[1][y][x] = piece.piece_type == chess.PAWN
			inputlayer[2][y][x] = piece.piece_type == chess.KNIGHT
			inputlayer[3][y][x] = piece.piece_type == chess.BISHOP or piece.piece_type == chess.QUEEN
			inputlayer[4][y][x] = piece.piece_type == chess.ROOK or piece.piece_type == chess.QUEEN
			inputlayer[5][y][x] = piece.piece_type == chess.KING

	return inputlayer

#print(rb.piece_at(0))
#print(rb)
#convert(rb) 	
#print(stockfish(rb, 15))


