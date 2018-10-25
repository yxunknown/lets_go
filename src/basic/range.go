package main

import "fmt"

/**
range 对slice或map进行迭代

index, value := range []T

*/

var power = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	for i, v := range power {
		fmt.Printf("%d:%d\n", i, v)
	}
	//omit the index
	for _, value := range power {
		if value%2 == 0 {
			fmt.Println(value, "is an even number")
		}
	}
	//omit the value
	for index := range power {
		power[index] = index*index + 2*index - 3
	}
	for i, v := range power {
		fmt.Printf("%d:%d\n", i, v)
	}
}
