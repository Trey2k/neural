package neural

import (
	"fmt"
	"os"
)

//Dot product
func Dot(v1, v2 []float64) float64 {
	if len(v1) == len(v2) {
		var result float64
		for i := 0; i < len(v1); i++ {
			result += v1[i] * v2[i]
		}
		return result
	}
	fmt.Println("Cannot do dot product on vectors not of the same size")
	os.Exit(1)
	return 0
}
