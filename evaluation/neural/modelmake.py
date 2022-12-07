import pandas
import numpy 
import chess

def board_convert(board):
	inputlayer = numpy.zeros((8, 8, 6), dtype=numpy.int8)

	for i in range(0,64):
		j = i
		if board.turn == chess.BLACK:
			j = i ^ 56

		piece = board.piece_at(i)
		y = j//8
		x = j%8
		if piece != None:
			inputlayer[y][x][0] = piece.color == board.turn
			inputlayer[y][x][1] = piece.piece_type == chess.PAWN
			inputlayer[y][x][2] = piece.piece_type == chess.KNIGHT
			inputlayer[y][x][3] = piece.piece_type == chess.BISHOP or piece.piece_type == chess.QUEEN
			inputlayer[y][x][4] = piece.piece_type == chess.ROOK or piece.piece_type == chess.QUEEN
			inputlayer[y][x][5] = piece.piece_type == chess.KING
	return inputlayer


def eval_convert(board, n):
	value = 0
	if n[0:2] == "#+":
		value = 1
	elif n[0:2] == "#-":
		value = -1
	else:
		value = numpy.tanh(float(n) / 500)
	
	if board.turn == chess.BLACK:
		value = -value
	return value

def inverse(board, n):
	if board.turn == chess.BLACK:
		n = -n
	return n


print("a")
f = pandas.read_csv('boards.csv')
print("b")

x1 = list(f['fen'])
y1 = list(f['eval'])
xs=[]
ys=[]

print("c")
# This step is taking way too long 
# TODO: figure out why

board = chess.Board()
for i in range(len(y1)):
	board.set_fen(x1[i])
	evaluation = y1[i]
	ys.append(eval_convert(board, evaluation))
	xs.append(board_convert(board))

print("d")

xs = numpy.array(xs)
ys = numpy.array(ys)
print("e")

import tensorflow
import tensorflow.keras.models as models
import tensorflow.keras.layers as layers
import tensorflow.keras.utils as utils
import tensorflow.keras.optimizers as optimizers
import tensorflow.keras.callbacks as callbacks

def build_model(conv_size, conv_depth):
	board3d = layers.Input(shape=(8, 8, 6), name="chessinput")
	x = board3d
	for i in range(conv_depth):
		x = layers.Conv2D(filters=conv_size, kernel_size=3, padding='same', activation='relu')(x)
	x = layers.Flatten()(x)
	x = layers.Dense(64, 'relu')(x)
	x = layers.Dense(1, 'tanh', name="chessoutput")(x)
	return models.Model(inputs=board3d, outputs=x)

model = build_model(16, 3)

model.compile(optimizer='adam',loss='mean_squared_error')
model.summary()

model.fit(xs,ys,epochs=100,batch_size=2048)
tensorflow.saved_model.save(model, "version8?")

# Version5 relu 3 conv layers, epochs = 10, 500k dataset
# Version6 ''', epochs = 25
# Version7 relu 3 conv layers new ordering,filters16,'''
# Version8 ''', epochs = 100