package main

import "fmt"

func main() {
	sum := 1
	//由于分号可以省略，因此for循环可以用做while循环的实现
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
