package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(newtonSqrt(4))
	fmt.Println(math.Sqrt(4))
}

func newtonSqrt(x float64) float64 {
	z := 1.0
	epsilon := 0.000001
	for {
		temp := z - ((z*z - x) / 2 * z)
		if math.Abs(z-temp) < epsilon {
			z = temp
			return z
		} else {
			z = temp
			fmt.Println(z)
		}
	}
}
