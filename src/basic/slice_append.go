package main

import "fmt"

/**
向slice添加元素
append([]T, T...)

len([]T) 计算一个数组的元素个数
cap([]T) 计算一个数组的容量
*/
func main() {
	var a []int //a is a nil slice now
	printSlices("a", a)

	a = append(a, 0) //a will grow as needed
	printSlices("a", a)

	a = append(a, 1)
	printSlices("a", a)

	a = append(a, 2, 3, 4, 5)
	printSlices("a", a)

}

func printSlices(s string, x []int) {
	fmt.Printf("%s len = %d cap=%d\n", s, len(x), cap(x))
}
