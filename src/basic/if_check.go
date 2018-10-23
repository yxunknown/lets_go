package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20))
	fmt.Println(
		max(2, 4),
		max(3, 2))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//if with a short statement
	//if 判断表达式前可以执行一个简单表达式，产生的变量只属于if这个域
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func max(a, b int) int {
	if v := a - b; v < 0 {
		//in if scope, v is available
		return a - v
	} else {
		//in if scope too, v is available
		return v + b
	}
	//out of if scope, v is not available
}
