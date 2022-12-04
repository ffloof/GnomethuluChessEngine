package neural

import (
	"github.com/galeone/tfgo"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	"fmt"
)

func Init(){
	model := tfgo.LoadModel("./evaluation/neural/output/keras/", []string{"serve"}, nil)

	fakeInput, _ := tf.NewTensor([1][6][8][8]float32{})
	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_chessinput", 0): fakeInput,
	})

	predictions := results[0]
	fmt.Println(predictions.Value())
}

