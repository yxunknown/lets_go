package main

import "fmt"

/**
if/else structure is like other languages, but
in go, you can declare variables before condition statements,
theses variables can be used on all branches
*/
func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// declare a variable num
	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		// num is usable in all branches
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
