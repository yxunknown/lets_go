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
	nums := []int{1, 2, 4, 2, 9, 19, 23}
	s := sum(nums...)
	fmt.Println(s)

	a := sum()
	fmt.Println(a)

	b := sumAtLeastOne(2)
	fmt.Println(b)

	// vararg function supports slice
	c := sumAtLeastOne(2, nums...)
	fmt.Println(c)
}

// this method can be called without any parameters
func sum(nums ...int) int {
	sum := 0
	for num := range nums {
		sum += num
	}
	return sum
}

func sumAtLeastOne(first int, rest ...int) int {
	var s = first
	for n := range rest {
		s += n
	}
	return s
}
