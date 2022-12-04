import pandas
import numpy 
import chess

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


def sigmoid(n): #Unrelated to neural net's sigmoid just for converting centipawn rating to a nice range	
	if n[0:2] == "#+":
		return 1
	elif n[0:2] == "#-":
		return 0
	else:
		return 1/(1+numpy.exp(-(float(n) / 100)))


f = pandas.read_csv('boards.csv')
xs = [convert(i) for i in list(f['fen'])]
ys = [sigmoid(i) for i in list(f['eval'])]
xs = numpy.array(xs)
ys = numpy.array(ys)
print(ys)


import tensorflow
import tensorflow.keras.models as models
import tensorflow.keras.layers as layers
import tensorflow.keras.utils as utils
import tensorflow.keras.optimizers as optimizers
import tensorflow.keras.callbacks as callbacks




def build_model(conv_size, conv_depth):
	board3d = layers.Input(shape=(6, 8, 8))
	x = board3d
	for _ in range(conv_depth):
		x = layers.Conv2D(filters=conv_size, kernel_size=3, padding='same', activation='relu')(x)
	x = layers.Flatten()(x)
	x = layers.Dense(64, 'relu')(x)
	x = layers.Dense(1, 'sigmoid')(x)
	return models.Model(inputs=board3d, outputs=x)

model = build_model(32, 4)

model.compile(optimizer='adam',loss= tensorflow.keras.losses.BinaryCrossentropy(from_logits=True),metrics=['accuracy'])

model.fit(xs,ys,epochs=100)
model.save("model.h5")