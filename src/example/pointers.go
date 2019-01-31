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
go supports pointers, allowing to pass references to values

& get memory address of variables(work on normal variables)
* read value from a pointer(work on pointers)
*Type pointers of Type

*/
func main() {
	a := 2
	fmt.Println("original", a)
	zeroval(a)
	fmt.Println("after zeroval", a)
	// &a: get memory address in hex of a, or a pointer to a somehow
	zeroptr(&a)
	fmt.Println("after zeroptr", a)

}

// pass value
func zeroval(ival int) {
	// ival is a copy of original variable
	// so any modifications on ival will not have influence on
	// original variable
	ival = 0
}

// pass reference
func zeroptr(iptr *int) {
	// modification value in memory addr
	// any references point to this memory addr will be influenced by
	// this modification
	*iptr = 0
}
