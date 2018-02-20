package special

import "math"

func Logistic(x float64) float64 {
	if x < 0 {
		z := math.Exp(x)
		return z / (1 + z)
	}
	return 1 / (1 + math.Exp(-x))
}
