package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	eps := 1e-9

	for d := (z*z - x) / (2 * z); math.Abs(d) > eps; d = (z*z - x) / (2 * z) {
		z -= d
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
