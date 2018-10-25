package main

import "fmt"

/**
fmt包定义的Stringer接口，类似于toString接口
type Stringer interface {
	String() string
}
*/

func main() {
	p := Person{"程飘", 24, "ChongQing"}
	fmt.Println(p)
}

type Person struct {
	name     string
	age      int
	location string
}

/**
toString() 的实现
*/
func (p Person) String() string {
	return fmt.Sprintf("Person[name: %v, age: %v, location: %v]",
		p.name, p.age, p.location)
}
