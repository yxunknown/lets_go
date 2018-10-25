package main

import "fmt"

func main() {
	m := make(map[string]int)

	//insert or modify an element
	m["one"] = 1
	//read value by key
	//if the key is not in the map, the map will return a zero value
	fmt.Println("the value", m["one"])
	//delete an element by key
	delete(m, "one")
	fmt.Println(m)
	//insert one
	m["two"] = 2

	//if key "two" is in map, element = map[key] and ok is true
	//otherwise element is zero value of value type, and ok is false
	element, ok := m["two"]
	fmt.Printf("Dose the map has the key? %v, the key is %v.\n", ok, element)

	three, isThreeExist := m["three"]
	fmt.Printf("Dose the map has the key? %v, the key is %v.\n", isThreeExist, three)
}
