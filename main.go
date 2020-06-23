package main

import (
	"fmt"
)

//Nurons slice of nurons
type Nurons []Nuron

//Nuron idk its a nuron man
type Nuron struct {
	Inputs   []float64
	Settings NuronSettings
}

//NuronSettingsSlice Slice
type NuronSettingsSlice []NuronSettings

//NuronSettings settings for nuron
type NuronSettings struct {
	Weights []float64
	Bias    float64
}

func main() {
	inputs := []float64{1, 2, 3, 2.5}
	weights := [][]float64{{0.2, 0.8, -0.5, 1.0}, {0.5, -0.91, 0.26, -0.5}, {-0.26, -0.27, 0.17, 0.87}}
	biases := []float64{2, 3, 0.5}

	settings := NewSettingsSlice(weights, biases)

	n := NewNurons(inputs, settings)

	fmt.Println(fmt.Sprint(CalcOutput(n)))

}

//CalcOutput calculate the output from a nuron
func (n *Nuron) CalcOutput() float64 {
	var output float64
	for i := 0; i < len(n.Inputs); i++ {
		output += n.Inputs[i] * n.Settings.Weights[i]
	}
	output += n.Settings.Bias
	return output
}

//CalcOutput calculate the output from a nuron
func CalcOutput(n Nurons) []float64 {
	var response []float64
	for i := 0; i < len(n); i++ {
		response = append(response, n[i].CalcOutput())
	}
	return response
}

//NewNuron make a New Nuron
func NewNuron(inputs []float64, settings NuronSettings) Nuron {
	var n Nuron
	n.Inputs = inputs
	n.Settings = settings
	return n
}

//NewNurons make a array of New Nurons
func NewNurons(inputs []float64, settings NuronSettingsSlice) Nurons {
	var n Nurons
	for i := 0; i < len(settings); i++ {
		n = append(n, NewNuron(inputs, settings[i]))
	}
	return n
}

//NewSettingsSlice stuff
func NewSettingsSlice(weights [][]float64, biases []float64) NuronSettingsSlice {
	var n NuronSettingsSlice
	if len(weights) == len(biases) {
		for i := 0; i < len(weights); i++ {
			n = append(n, NuronSettings{weights[i], biases[i]})
		}
	}
	return n
}
