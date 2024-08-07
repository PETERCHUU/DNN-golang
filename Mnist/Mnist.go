package Mnist

import (
	"fmt"

	"github.com/PETERCHUU/DNNGolang"
	"github.com/PETERCHUU/DNNGolang/Mnist/FileReader"
	"github.com/PETERCHUU/DNNGolang/function"
)

const (
	trainingDataPath  = "Mnist/train/train-images.idx3-ubyte"
	trainingLabelPath = "Mnist/train/train-labels.idx1-ubyte"
	testDataPath      = "Mnist/train/t10k-images.idx3-ubyte"
	testLabelPath     = "Mnist/train/t10k-labels.idx1-ubyte"

	sampleRate           = 1000
	learningRate float64 = 0.15
)

func Run() DNNGolang.Chain {
	module := DNNGolang.NewNetwork().FCLayer(784, 49, function.Sigmoid, learningRate).
		FCLayer(49, 10, function.Softmax, learningRate)
	betterModule := module.Copy()
	var accurate float64
	sample := FileReader.InitSample(trainingDataPath, trainingLabelPath)
	tester := FileReader.InitSample(testDataPath, testLabelPath)

	fmt.Printf("Accurate before train: %.2f\n", CalculateAccurate(&module, tester))

	// for i, v := range sample {
	// 	module.BackProp(v.Image[:], v.Label[:], learningRate)
	// 	//fmt.Printf("before weight %.2f", (*(*module.Layers)[2].Neurons)[3].Weights)
	// 	if i%1000 == 0 {
	// 		fmt.Printf("Accurate after %d train: %.4f\n", i, calculateAccurate(&module, tester))
	// 	}
	// }

	for i := 0; i < len(sample); i += sampleRate {
		sampleInput := make([][]float64, sampleRate)
		sampleTarget := make([][]float64, sampleRate)
		for j := 0; j < sampleRate; j++ {
			sampleInput[j] = sample[i+j].Image[:]
			sampleTarget[j] = sample[i+j].Label[:]
		}

		module.UpdateMiniBatch(sampleInput, sampleTarget, sampleRate, learningRate)
		thisAccurate := CalculateAccurate(&module, tester)
		if thisAccurate > accurate {
			accurate = thisAccurate
			betterModule = module.Copy()
		}
		fmt.Printf("Accurate after %d train: %.4f\n", i, thisAccurate)
	}

	//fmt.Printf("after weight %.2f", (*(*betterModule.Layers)[2].Neurons)[3].Weights)

	accurate = CalculateAccurate(&betterModule, tester)

	fmt.Printf("Accurate after train: %.4f\n", accurate)
	return betterModule
}

func CalculateAccurate(module *DNNGolang.Chain, sample []FileReader.MnstSample) float64 {
	var accurate float64
	for _, v := range sample {
		predict := module.Predict(v.Image[:])
		accurate += DNNGolang.Accurate(predict, v.Label[:])
	}

	accurate /= float64(len(sample))
	if accurate > 1 {
		accurate -= 1
	}
	return accurate
}
