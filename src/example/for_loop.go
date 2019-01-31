package main

import "fmt"

/**
go has only for loop structure
for without condition statements is a dead-loop
*/

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("a")
		fmt.Println("b")
		fmt.Println("c")
		fmt.Println("d")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
