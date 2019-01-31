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
range is used to iterates
*/

func main() {
	nums := []int{1, 2, 3}
	// i is index, v is value
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// range used on map
	kvs := map[string]string{"a": "112", "b": "121", "c": "114"}
	// k is key, and v is value
	for k, v := range kvs {
		fmt.Println(k, "is map to", v)
	}

	// range used on string
	// i is index, and c is unicode code points
	for i, c := range "china" {
		fmt.Println(i, c)
	}
}
