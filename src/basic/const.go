package main

import (
	"fmt"
	"math"
)

/**
const 用于定义常量，与普通变量定义类似
但是常量可以使用 := 语法
*/

const Pi = math.Pi

func main() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")
	const truth = true
	fmt.Println("Go rules?", truth)
}
