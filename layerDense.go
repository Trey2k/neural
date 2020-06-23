package neural

import "math/rand"

//LayerDense stuff
type LayerDense struct {
	Nurons     Nurons
	NuronCount int
}

//NewLayerDense stuff
func NewLayerDense(inputCount int, nuronCount int) *LayerDense {
	rand.Seed(0)
	var layer LayerDense
	layer.Nurons = make(Nurons, nuronCount)
	for x := 0; x < nuronCount; x++ {
		for i := 0; i < inputCount; i++ {
			layer.Nurons[x].Settings.Weights = append(layer.Nurons[x].Settings.Weights, 0.10*rand.Float64())
		}
		layer.Nurons[x].Settings.Bias = rand.Float64()
	}
	return &layer
}

//CalculateOutput calculate the output from a nuron
func (layer *LayerDense) CalculateOutput(input Batch) Batch {
	return CalculateOutput(layer.Nurons, input)
}
