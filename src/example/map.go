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
maps are key-value type, and map is disordered type
*/
func main() {
	// create empty map
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["zero"] = 0
	fmt.Println(m)
	one := m["one"]
	fmt.Println(one)
	// the second return value is to mark the key in or not in the map
	// if the key not in the map, will return a zero-value of value-type
	// and exits is false
	v1, exits := m["three"]
	fmt.Println(v1, exits)

	// delete  a key in map
	delete(m, "three")
	fmt.Println(m)

	// liberal
	// declaration and initialization
	numToStr := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(numToStr)
}
