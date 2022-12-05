import pandas
import numpy 
import chess

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


def sigmoid(n): #Unrelated to neural net's sigmoid just for converting centipawn rating to a nice range	
	if n[0:2] == "#+":
		return 1
	elif n[0:2] == "#-":
		return -1
	else:
		return numpy.tanh(float(n) / 350)

def inverse(fen, n):
	board = chess.Board()
	board.set_fen(fen)

	if board.turn == chess.BLACK:
		n = -n
	return n


f = pandas.read_csv('boards.csv')

xs = list(f['fen'])
ys = list(f['eval'])

for i in range(len(ys)):
	ys[i] = sigmoid(ys[i])
	ys[i] = inverse(xs[i], ys[i])
	xs[i] = convert(xs[i])


xs = numpy.array(xs)
ys = numpy.array(ys)

import tensorflow
import tensorflow.keras.models as models
import tensorflow.keras.layers as layers
import tensorflow.keras.utils as utils
import tensorflow.keras.optimizers as optimizers
import tensorflow.keras.callbacks as callbacks

def build_model(conv_size, conv_depth):
	board3d = layers.Input(shape=(6, 8, 8), name="chessinput")
	x = board3d
	for i in range(conv_depth):
		x = layers.Conv2D(filters=conv_size, kernel_size=3, padding='same', activation='sigmoid')(x)
	x = layers.Flatten()(x)
	x = layers.Dense(64, 'sigmoid')(x)
	x = layers.Dense(1, 'tanh', name="chessoutput")(x)
	return models.Model(inputs=board3d, outputs=x)

model = build_model(32, 3)

model.compile(optimizer='adam',loss='mean_squared_error')
model.summary()

model.fit(xs,ys,epochs=20)
tensorflow.saved_model.save(model, "version1")