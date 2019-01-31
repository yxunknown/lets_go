package main

import (
	"fmt"
	"math"
)

/**
go supports constants for character, string, boolean
and numeric values.

numeric constants has no type until it's given one, such as by
an explict cast
*/

// global constant
const s = "string"

func main() {
	const n = 50000 // n is untype int now
	fmt.Println(s, n)

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n)) // n is explictly casted to float64
}
