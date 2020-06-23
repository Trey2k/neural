package main

import (
	"fmt"
)

//Nurons slice of nurons
type Nurons []Nuron

//Nuron idk its a nuron man
type Nuron struct {
	Inputs []Input
	Bias   float64
}

//Input input variables for nurons
type Input struct {
	Input  float64
	Weight float64
}

func main() {
	inputs := []float64{1, 2, 3, 2.5}
	weights1 := []float64{0.2, 0.8, -0.5, 1.0}
	weights2 := []float64{0.5, -0.91, 0.26, -0.5}
	weights3 := []float64{-0.26, -0.27, 0.17, 0.87}
	n := Nurons{
		NewNuron(inputs, weights1, 2),
		NewNuron(inputs, weights2, 3),
		NewNuron(inputs, weights3, 0.5),
	}

	fmt.Println(fmt.Sprint(CalcOutput(n)))

}

//CalcOutput calculate the output from a nuron
func (n *Nuron) CalcOutput() float64 {
	var output float64
	for i := 0; i < len(n.Inputs); i++ {
		output += n.Inputs[i].Input * n.Inputs[i].Weight
	}
	output += n.Bias
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
func NewNuron(inputs []float64, weights []float64, bias float64) Nuron {
	var n Nuron

	for i := 0; i < len(inputs); i++ {
		n.Inputs = append(n.Inputs, Input{
			Input:  inputs[i],
			Weight: weights[i],
		})
	}

	n.Bias = bias

	return n
}
