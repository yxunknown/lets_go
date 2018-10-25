package main

import (
	"fmt"
	"math"
)

func main() {
	//function is a type
	//function can be assigned to a variable
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}
