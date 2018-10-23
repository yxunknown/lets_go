package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 3}
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)
	//the fields of struct
	v.x = 4
	v.y = 5
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)
	//the pointer of struct
	p := &v
	fmt.Println(*p)
	//access filed via pointer is transparent
	fmt.Println((*p).x)
	fmt.Println(p.y)
}
