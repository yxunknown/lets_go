package main

import "fmt"

/**
slice 指向一个序列的值，并包含其长度信息
*/
func main() {
	p := []int{2, 3, 5, 7, 9, 11}
	fmt.Println("p ==", p)
	const f = "p[%d] == %d\n"
	for i := 0; i < len(p); i++ {
		fmt.Printf(f, i, p[i])
	}
}
