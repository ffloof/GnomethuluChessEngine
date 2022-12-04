import numpy
import chess
from tensorflow.keras import models

model = models.load_model('modelv2.h5')

def convert(board_fen):
	board = chess.Board()
	board.set_fen(board_fen)
	inputlayer = numpy.zeros((6, 8, 8), dtype=numpy.int8)

	for i in range(0,64):
		piece = board.piece_at(i)
		y = i//8
		x = i%8
		if piece != None:
			inputlayer[0][y][x] = piece.color == board.turn
			inputlayer[1][y][x] = piece.piece_type == chess.PAWN
			inputlayer[2][y][x] = piece.piece_type == chess.KNIGHT
			inputlayer[3][y][x] = piece.piece_type == chess.BISHOP or piece.piece_type == chess.QUEEN
			inputlayer[4][y][x] = piece.piece_type == chess.ROOK or piece.piece_type == chess.QUEEN
			inputlayer[5][y][x] = piece.piece_type == chess.KING
	return inputlayer

board3d = convert("r2q1rk1/p3ppbp/1pn2np1/2pp4/3P1Bb1/P1PBPN2/1P1N1PPP/R2Q1RK1 w - - 1 11")
board3d = numpy.expand_dims(board3d, 0)
print(model(board3d)[0][0])