package nnfcgolang

func (c Chain) calculateAccurate(inputs, targets [][]float32) float32 {
	var accurate float32
	if len(inputs) != len(targets) {
		panic("dataFormate error, input len != target len")
	}
	for i := range inputs {
		predict := c.Predict(inputs[i])

		accurate += Accurate(predict, targets[i])
	}
	accurate /= float32(len(inputs))
	return accurate
}
