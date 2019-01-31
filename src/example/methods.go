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

/**
methods are defined on structs
receiver可以定义为值类型或者指针类型，在进行使用的时候，
go会进行自动转换。
pointers type receiver use to avoid copying on method calls
and allow to the method to mutate the receiving struct
*/

func main() {
	r := Rect{10, 24}
	areasOfR := r.area()
	fmt.Println(areasOfR)
	fmt.Println(r)
	r.adjustToSquare()
	fmt.Println(r)
}

type Rect struct {
	width, height int
}

// value type receiver
// just read value of struct, no modification
func (r Rect) area() int {
	return r.height * r.width
}

// pointer type receiver
// can read value of struct, and do modifications
func (r *Rect) adjustToSquare() {
	if r.width > r.height {
		r.height = r.width
	} else {
		r.width = r.height
	}
}
