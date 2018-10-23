package main

import "fmt"

func main() {
	v := 3
	f := 3.14
	c := 0.993 + 1.3223i
	s := "strings"
	const format = "%v is %T type\n"
	fmt.Printf(format, v, v)
	fmt.Printf(format, f, f)
	fmt.Printf(format, c, c)
	fmt.Printf(format, s, s)

}
