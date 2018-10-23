package main

import "fmt"

/**
defer用于延迟当前语句的执行
直到上层函数返回后，再执行当前语句
defer后的参数是立即生成的，但直到上层函数返回后才会被执行
*/
func main() {
	//fmt.Println("World") will be executed after the main
	//func returned
	defer fmt.Println("World")
	fmt.Println("Hello")
}
