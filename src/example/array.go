package main

import "fmt"

/**
in go, type and length are part of an array
arrays are accessed by index
arrays are value-type, when you assign or pass around an array value,
you will make a copy of its contents. to avoid this copy operation,
can use pointer.
*/
func main() {
	// a is an int array with length is 5
	// a is initialized with zero-value [0 0 0 0 0]
	var a [5]int
	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("set a[4]", a)
	fmt.Println("get a[4]", a[4])

	// len is a built-in function to calculate length of an array type
	fmt.Println("length of a", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// make the compiler to count the elements for array
	c := [...]string{"渝中", "南岸", "江北"}
	fmt.Println(c, len(c))

	var twoD [2][2]int
	fmt.Println(twoD)

	for i := 1; i <= len(twoD); i++ {
		for j := 1; j <= len(twoD[i-1]); j++ {
			twoD[i-1][j-1] = j
		}
	}
	fmt.Println(twoD)
}
