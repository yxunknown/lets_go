package main

import "fmt"

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i))
	}
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		//here can access variables out of this function
		//函数闭包和所访问的变量被绑定在一起
		sum += x
		return sum
	}
}
