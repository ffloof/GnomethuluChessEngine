package neural

import (
	"github.com/galeone/tfgo"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	"github.com/dylhunn/dragontoothmg"
)

func convert(board dragontoothmg.Board) *tf.Tensor {
	inputs := [1][8][8][6]float32{}

	for i:=0;i<64;i++ {
		piece, pieceColor := dragontoothmg.GetPieceType(uint8(i),&board)

		j := i
		if !board.Wtomove {
			j = i ^ 56
		}

		y := j/8
		x := j%8
		if piece != 0 {
			if pieceColor == board.Wtomove { inputs[0][y][x][0] = 1 }
			if piece == dragontoothmg.Pawn { inputs[0][y][x][1] = 1 } 
			if piece == dragontoothmg.Knight { inputs[0][y][x][2] = 1 }
			if piece == dragontoothmg.Bishop || piece == dragontoothmg.Queen { inputs[0][y][x][3] = 1 } 
			if piece == dragontoothmg.Rook || piece == dragontoothmg.Queen { inputs[0][y][x][4] = 1} 
			if piece == dragontoothmg.King { inputs[0][y][x][5] = 1 }
		}
	}

	inputs2, _ := tf.NewTensor(inputs)
	return inputs2
}



func Init() func (board dragontoothmg.Board) float64 {
	model := tfgo.LoadModel("./evaluation/neural/version6/", []string{"serve"}, nil)

	return func (board dragontoothmg.Board) float64 {
		boardInput := convert(board)
		results := model.Exec([]tf.Output{
			model.Op("StatefulPartitionedCall", 0),
		}, map[tf.Output]*tf.Tensor{
			model.Op("serving_default_chessinput", 0): boardInput,
		})

		predictions := results[0]
		return float64((predictions.Value().([][]float32)[0][0]))
	}
}

