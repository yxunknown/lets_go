package main

import "fmt"

func main() {
	fmt.Println("hello world")
	s := []string{"a"}
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = append(s, "2")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = append(s, "4", "5", "6")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
