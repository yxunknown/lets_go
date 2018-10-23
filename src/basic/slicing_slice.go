package main

import "fmt"

/**
可以对slice进行切片
表达式
slice[low:high] return slice[low]...slice[high - 1]
slice[low:low] return []
slice[low:low + 1] return slice[low]
*/
func main() {
	p := []int{2, 3, 5, 7, 9, 11}
	fmt.Println(p)
	fmt.Println("p[1:2] is", p[1:2])
	fmt.Println("p[0:4] is", p[0:4])
	//从开始到n
	fmt.Println("p[:5] is", p[:5])
	//从n到结束
	fmt.Println("p[2:] is", p[2:])
}
