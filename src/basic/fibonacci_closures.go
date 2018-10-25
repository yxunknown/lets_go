package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//implement fibonacci array by function closure
//0 1 1 2 3 5 8 13 21 34
func fibonacci() func() int {
	pre := 0
	last := 1
	return func() int {
		//i don't know what happens here
		//but this function gets work
		fibonacci := pre
		pre = last
		last = fibonacci + pre
		return fibonacci
	}
}
