package main

import "fmt"

type Verte struct {
	x int
	y int
}

var (
	v1 = Verte{1, 2}  // type of Vertex
	v2 = Verte{x: 4}  // y = 0 is omit
	v3 = Verte{y: 3}  // x = o is omit
	p  = &Verte{1, 3} // pointer *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, *p)
}
