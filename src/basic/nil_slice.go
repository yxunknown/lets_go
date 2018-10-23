package main

import "fmt"

/**
slice的长度和容量都为0时，该slice为 nil
*/
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
