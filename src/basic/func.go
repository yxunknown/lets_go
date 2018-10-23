package main

import "fmt"

func main() {
	fmt.Println("23 + 31 = ", add(23, 31))
	fmt.Println("23 - 12 = ", minus(23, 12))
}

func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}
