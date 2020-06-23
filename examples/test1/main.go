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

	settings1 := neural.NuronSettingsSlice{
		neural.NuronSettings{[]float64{0.2, 0.8, -0.5, 1.0}, 1},
		neural.NuronSettings{[]float64{0.5, -0.91, 0.26, -0.5}, 2},
		neural.NuronSettings{[]float64{-0.26, -0.27, 0.17, 0.87}, 0.5},
	}

	settings2 := neural.NuronSettingsSlice{
		neural.NuronSettings{[]float64{0.1, -0.14, 0.5}, -1},
		neural.NuronSettings{[]float64{-0.5, 0.12, -0.33}, 2},
		neural.NuronSettings{[]float64{-0.44, 0.73, -0.13}, -0.5},
	}

	layer1Outputs := neural.CalculateOutput(neural.NewNurons(settings1), X)
	layer2Outputs := neural.CalculateOutput(neural.NewNurons(settings2), layer1Outputs)

	fmt.Println(fmt.Sprint(layer2Outputs))

}
