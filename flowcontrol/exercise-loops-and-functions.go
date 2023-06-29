package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 1e-9

	for d := (z*z - x) / (2 * z); math.Abs(d) > eps; d = (z*z - x) / (2 * z) {
		z -= d
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
