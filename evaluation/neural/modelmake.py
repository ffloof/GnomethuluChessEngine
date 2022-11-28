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

model.compile(optimizer='adam',loss= tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True),metrics=['accuracy'])
