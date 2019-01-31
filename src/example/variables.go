package main

import "fmt"

/**
go will infer type info of variables, you can declare multiple variables at once.
variables declared without corresponding initialization are zero-valued.
zero-value of each type:
number 0
boolean false
string "" (blank string)
complex (0+0i)
*/
func main() {
	// declare a variable named a and give it an initialization value
	// go infers type of a is string
	var a = "initial"
	fmt.Println(a)

	// declaration multiple variables at once
	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// declare a variable without initialization
	// e will be given a zero-value by compiler
	var e int
	fmt.Println(e)

	// short declaration
	f := "short"
	fmt.Println(f)

	// multiple variables with short declaration
	g, h := "g", "h"
	fmt.Println(g, h)
}
