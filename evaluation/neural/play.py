import numpy
import chess
from tensorflow.keras import models

model = models.load_model('modelv4.h5')

def convert(board_fen):
	board = chess.Board()
	board.set_fen(board_fen)
	inputlayer = numpy.zeros((6, 8, 8), dtype=numpy.int8)

	for i in range(0,64):
		j = i
		if board.turn == chess.BLACK:
			j = i ^ 56

		piece = board.piece_at(i)
		y = j//8
		x = j%8
		if piece != None:
			inputlayer[0][y][x] = piece.color == board.turn
			inputlayer[1][y][x] = piece.piece_type == chess.PAWN
			inputlayer[2][y][x] = piece.piece_type == chess.KNIGHT
			inputlayer[3][y][x] = piece.piece_type == chess.BISHOP or piece.piece_type == chess.QUEEN
			inputlayer[4][y][x] = piece.piece_type == chess.ROOK or piece.piece_type == chess.QUEEN
			inputlayer[5][y][x] = piece.piece_type == chess.KING
	return inputlayer

board3d = convert("r4rk1/ppp2ppp/2nb4/1B6/8/2P1PN2/bB3PPP/5RK1 b - - 1 14")
board3d = numpy.expand_dims(board3d, 0)
print(model(board3d)[0][0])