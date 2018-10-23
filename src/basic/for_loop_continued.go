package main

import "fmt"

/**
前置、后置语句为空
*/
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
