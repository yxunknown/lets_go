package main

import "fmt"

func main() {
	i, j := 2, 2701
	p := &i         //point to i
	fmt.Println(*p) //read value via pointer
	*p = 21         //set value via pointer
	fmt.Println(i)  //value of i has been changed

	p = &j //point to j
	//pointer can be reassigned
	*p = *p / 37 //use value via pointer
	//unlike c, the pointers don't support for operations in Go
	fmt.Println(j)
}
