package main

import "fmt"

/**
when declare an variable but not initialize it
this variable will be initialized a zero value
number to be 0
bool to be false
string to be ""
*/
func main() {
	var i int
	var f float64
	var b bool
	var s string
	var c complex128
	fmt.Println(i, f, b, s, c)
}
