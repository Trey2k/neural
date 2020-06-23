package neural

//Nurons slice of nurons
type Nurons []Nuron

//Nuron idk its a nuron man
type Nuron struct {
	Settings NuronSettings
}

type Batch struct {
	Batch [][]float64
}

//NuronSettingsSlice Slice
type NuronSettingsSlice []NuronSettings

//NuronSettings settings for nuron
type NuronSettings struct {
	Weights []float64
	Bias    float64
}

//CalculateOutput calculate the output from a nuron
func (n *Nuron) CalculateOutput(input Batch) Batch {
	var output Batch
	output.Batch = make([][]float64, len(input.Batch))
	for i := 0; i < len(input.Batch); i++ {
		output.Batch[i] = append(output.Batch[i], Dot(n.Settings.Weights, input.Batch[i])+n.Settings.Bias)
	}
	return output
}

//CalculateOutput calculate the output from a nuron
func CalculateOutput(n Nurons, input Batch) Batch {
	var output Batch
	output.Batch = make([][]float64, len(input.Batch))
	for i := 0; i < len(input.Batch); i++ {
		for x := 0; x < len(n); x++ {
			output.Batch[i] = append(output.Batch[i], Dot(n[x].Settings.Weights, input.Batch[i])+n[x].Settings.Bias)
		}
	}
	return output
}

//NewNuron make a New Nuron
func NewNuron(settings NuronSettings) Nuron {
	var n Nuron
	n.Settings = settings
	return n
}

//NewNurons make a array of New Nurons
func NewNurons(settings NuronSettingsSlice) Nurons {
	var n Nurons
	for i := 0; i < len(settings); i++ {
		n = append(n, NewNuron(settings[i]))
	}
	return n
}
