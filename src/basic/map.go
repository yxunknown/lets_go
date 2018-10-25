package main

import "fmt"

/**
map需要使用make来创建
*/
var pmToVendor map[string]string

func main() {
	pmToVendor = make(map[string]string)
	pmToVendor["china mobile"] = "13232323312"
	fmt.Println(pmToVendor)
	fmt.Println(pmToVendor["china mobile"])
}
