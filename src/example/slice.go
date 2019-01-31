/**
 *
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
	"fmt"
)

/**
unlike array, slice are typed by the elements it contains,
so slice can change its length.
slice is built on array, but support more powerful operations and conveniences

structure of slice

++++++++++++++
+ ptr  *Elem +
++++++++++++++
+ len  int   +
++++++++++++++
+ cap  int   +
++++++++++++++

when doing slice operations on slice, these operations dont make a copy of slice,
but to create a new slice and use a pointer to point the original array in that slice,
so any modifications on new slice will pass to the original slice.

slice can't not been grown beyond its capacity.
but can use append or copy function to create new slice

resizing slice does'nt make a copy of underlying array, the full array will kept
in memory until it is no longer referenced. this may cause the program hold all
the data in memory but only use a small piece of it is needed.
*/
func main() {
	// create a slice with string type
	s := make([]string, 3)
	fmt.Println("emp", s)
	s[0] = "ch"
	s[1] = "i"
	s[2] = "na"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	// append elements into slice
	s = append(s, "is")
	s = append(s, "a", "big", "country")
	fmt.Println("append", s)

	c := make([]string, len(s))
	// copy slice from another slice
	copy(c, s)
	fmt.Println("copy", c)

	// slice operation
	// slicing is sharing the same storage as original slice
	// start at 2 and end at 4, return 3 elements in total
	l := s[2:5]
	fmt.Println(l)

	// end at 4
	// start at 0 by default
	l = s[:5]
	fmt.Println(l)

	// start at 2
	// end at len of slice by default
	l = s[2:]
	fmt.Println(l)

	// l is equal to s
	l = s[:]
	fmt.Println(l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = j + 1
		}
	}
	fmt.Println("2d", twoD)

	// let the compiler to count elements
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	// slice can created by calling make function
	// make([]t, len, cap) []T
	var bs []byte
	bs = make([]byte, 5, 5)
	fmt.Println("bs", bs)
	bs = append(bs, 2)
	fmt.Println("bs", bs)
	fmt.Println("len", len(bs))
	fmt.Println("cap", cap(bs))
	// cap of a slice will auto-increase to fit in the elements in slice
	// cap greater or equal than len

	// the zero value of a slice is nil, the len and cap functions will both return
	// 0 for a nil slice
	var nils []int
	fmt.Println("nil len", len(nils))
	fmt.Println("nil cap", cap(nils))

	m := make([]int, 5, 5)
	for i := 0; i < len(m); i++ {
		m[i] = i + 1
	}
	// create new slice
	n := make([]int, len(m), (cap(m)+1)*2)
	// copy all values from m to n
	for i := range n {
		n[i] = m[i]
	}
	// assign n to m
	m = n

	// copy function implements better than for loop
	x := make([]int, len(m), (cap(m)+1)*2)
	copy(x, m)
	m = x

	p := []int{1, 2, 3}
	p = appendIntSlice(p, 4, 5, 6, 7)
	fmt.Println(p)

	// append function
	f := make([]int, 1)
	f = append(f, 1, 2, 3)
	fmt.Println(f)
	// append a slice to another
	s1 := []string{"china", "is"}
	s2 := []string{"a", "big", "country"}
	s1 = append(s1, s2...) // equal to append(s1, s2[0], s2[1], s2[2])

	xs := s2[1:3]
	xs[0] = "super"
	fmt.Println(xs)
	fmt.Println(s2)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// pass function as parameter
	evenNums := filter(nums, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(nums)
	fmt.Println(evenNums)

}

func appendIntSlice(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		t := make([]int, (n+1)*2)
		copy(t, slice)
		slice = t
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// accept a function as parameter
func filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
