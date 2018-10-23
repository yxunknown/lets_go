package main

import "fmt"

/**
slice 由 make 函数创建， make函数返回一个零长度的数组并返回给一个slice只想这个数组

a := make([]int, 5) // len(a) = 5

make可以限定一个数组的最大容量
make([]int, 0, 5)
*/
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap=%d\n", s, len(x), cap(x))
}
