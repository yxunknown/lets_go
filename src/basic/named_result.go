package main

import "fmt"

func main() {
	x, y := swapint(124)
	fmt.Println("x is", x, "y is", y)
}

func swapint(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
