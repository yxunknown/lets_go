/*
 * *
 *    Copyright [${YEAR}] ${USER}
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package main

import "fmt"

func main() {
	p := Person{"mevur", 24}
	fmt.Println(p)
	p = Person{name: "hercats", age: 23}
	fmt.Println(p)
	p = Person{name: "sekai"}
	// when initializing a struct, omitted fields will be zero-value
	fmt.Println(p)

	p = Person{"FLASH", 25}
	ptr := &p
	fmt.Println(ptr)
	fmt.Println(ptr.name, ptr.age)

	// struct is mutable type
	ptr.name = "flash"
	fmt.Println(p, ptr)
	ptr.age = 23
	fmt.Println(p, ptr)
}

type Person struct {
	name string
	age  int
}
