package main

import "fmt"

func main() {
	var i, j, k int = 2, 3, 5
	m := 2 // this declaration can not use out of function
	c, python, java := true, false, "nope!"
	fmt.Println(i, j, k, m, c, python, java)
}
