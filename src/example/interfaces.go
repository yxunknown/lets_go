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

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/**
interface is two things: a set of methods, or a type

the interface{} type, 一个没有方法的接口，由于go没有显示声明实现接口
的操作，而任何类型都有0个或多个方法，因此任何类型都满足 interface{} 接口。

因此，如果一个方法接受的参数类型为 interface{}， 则表示所有的类型都可以
传入该方法进行处理。

*/

func main() {
	animals := []Animal{&Dog{}, new(Cat), Llama{}, PythonProgramer{}}

	for _, animal := range animals {
		fmt.Println(animal.speak())
	}

	ele := make([]interface{}, len(animals))
	for i, v := range animals {
		ele[i] = v
	}
	printAll(ele)

	printEvery(ele...)

	// parse json data
	var val map[string]Timestamp
	var input = `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012"
	}`
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}
	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}

type Timestamp time.Time

// implement the method that convert from json
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	// t is a copy of a another pointer,
	// these to pointers are pointing to a same one value
	// if we want that any modifications happens here can be seen
	// by caller, we can dereference the pointer and modify the value
	// of the pointer directly.
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	// convert a time.Time type to Timestamp
	*t = Timestamp(v)
	return nil
}

type Animal interface {
	speak() string
}

// any types that define speak method is
// satisfy the Animal interface.
// there is no implements keyword in go, so whether or not a
// type satisfies an interface is determined automatically.

type Dog struct{}

func (d Dog) speak() string {
	return "woof!"
}

type Cat struct{}

func (c *Cat) speak() string {
	return "meow!"
}

type Llama struct{}

func (l Llama) speak() string {
	return "????"
}

type PythonProgramer struct{}

func (pp PythonProgramer) speak() string {
	return "oh fxxk, what type of this!"
}

func some(v interface{}) {
	// what is the type of v?
	// v is not of any type, but is of interface{} type
	// when passing a variable into this function,
	// go converts type of v into interface{}
}

//how to use []interface{}
func printAll(eles []interface{}) {
	for _, ele := range eles {
		fmt.Printf("%v ", ele)
	}
	fmt.Println()
}

func printEvery(eles ...interface{}) {
	for _, v := range eles {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}
