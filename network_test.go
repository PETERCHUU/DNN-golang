package nnfcgolang_test

import (
	"fmt"
	"nnfcgolang"
	"nnfcgolang/function"
	"testing"
)

const learningRate = 0.1

func TestNewNetwork(T *testing.T) {
	FCnetwork := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate).
		FCLayer(8, 4, function.Sigmoid, learningRate).FCLayer(4, 1, function.Softmax, learningRate)

	for i, v := range *FCnetwork.Layers {
		for j, n := range *v.Neurons {
			T.Log(i, j, n)
		}
	}
	//FCnetwork.trainBy(backprop.adam)
	//FCnetwork.windowsize(int)
	//FCnetwork.train(data)
	//FCnetwork.test(testdata)
	//newwork.saveAs("model.toml")
}

func BenchmarkNewNetwork(T *testing.B) {
	FCnetwork := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate).
		FCLayer(8, 4, function.Sigmoid, learningRate).FCLayer(4, 1, function.Softmax, learningRate)
	for i, v := range *FCnetwork.Layers {
		for j, n := range *v.Neurons {
			fmt.Println(i, j, n)
		}
	}
	//FCnetwork.trainBy(backprop.adam)
	//FCnetwork.windowsize(int)
	//FCnetwork.train(data)
	//FCnetwork.test(testdata)
	//newwork.saveAs("model.toml")
}

func TestFCLayer(t *testing.T) {
	nn := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate)
	println(len(*nn.Layers))
	t.Log("haha")
}

func TestPredict(t *testing.T) {
	FCnetwork := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate).
		FCLayer(8, 4, function.Sigmoid, learningRate).FCLayer(4, 1, function.Softmax, learningRate)
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	output := FCnetwork.Predict(data)
	fmt.Println()
	for i := range output {
		t.Log(output[i])
	}
}

func TestRNNpredict(t *testing.T) {
	nn := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate).
		FCLayer(8, 4, function.Sigmoid, learningRate).FCLayer(4, 1, function.Softmax, learningRate)
	nn.Predict([]float32{1, 2, 3, 4, 5, 6, 7, 8})

}

func BenchmarkPredict(t *testing.B) {
	FCnetwork := nnfcgolang.NewNetwork().FCLayer(16, 8, function.Sigmoid, learningRate).
		FCLayer(8, 4, function.Sigmoid, learningRate).FCLayer(4, 1, function.Softmax, learningRate)
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	output := FCnetwork.Predict(data)
	fmt.Println()
	for i := range output {
		t.Log(output[i])
	}
}
