package main

import (
	"fmt"

	"github.com/Trey2k/neural"
)

func main() {
	X := neural.Batch{[][]float64{
		{1, 2, 3, 2.5},
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}}

	layer1 := neural.NewLayerDense(4, 5)
	layer2 := neural.NewLayerDense(5, 2)

	layer1Output := layer1.CalculateOutput(X)
	layer2Output := layer2.CalculateOutput(layer1Output)

	fmt.Println(fmt.Sprint(layer2Output))

}
