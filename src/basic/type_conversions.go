package main

import (
	"fmt"
	"math"
)

/**
type-cast expression
type(value)

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

short variable declaration
i := 42
f := float64(i)
u := uint(f)
*/
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, f, z)
}
