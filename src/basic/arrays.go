package main

import "fmt"

/**
definition of array
[n]T

var a [10]int // an array of int
*/
func main() {

	var a [10]string // the element is initialized with zero value
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
}
