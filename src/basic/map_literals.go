package main

import "fmt"

var m = map[string]int{
	//在map的声明中写明类型后，元素的创建可以省略类型名
	"zero": 0, "one": 1, "two": 2,
}

func main() {
	fmt.Println(m)
}
