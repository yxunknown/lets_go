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
	"errors"
	"fmt"
)

/**
in go way, usually return error value at the last returned value,
which type is error.

errors.New return a errorString error：
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	// return a pointer
    return &errorString{text}
}

fmt package formats an error value by calling its Error method

fmt.Errorf is used to produce an informative error info
return an errorString error
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
*/

func main() {
	for _, i := range []int{3, 4, 42, 23, 35} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{3, 4, 42, 23, 35} {
		if r, e := f2(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f2(42)
	// handler error type
	// try to convert to the specified type
	if argErr, convertOk := e.(*ArgError); convertOk {
		fmt.Println("encounter a ArgError:", argErr)
	} else {
		fmt.Println("other error occurs")
	}

	if _, e := f3(42); e != nil {
		fmt.Println(e)
	}

}

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors is to make some error with text ms
		return -1, errors.New("can't work with 42")
	}
	// use nil to indicate there is no error
	return arg + 3, nil
}

// custom error and implement error interface
type ArgError struct {
	arg int
	msg string
}

// base on officer guide, receiver should use pointer
func (arg *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", arg.arg, arg.msg)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// pointer interface can be used only by pointer
		return -1, &ArgError{arg: arg, msg: "can't work with it"}
	}
	return arg + 3, nil
}

// use fmt.Errorf
func f3(arg int) (int, error) {
	if arg == 42 {
		return -1, fmt.Errorf("%d - can't work with it", arg)
	}
	return arg + 3, nil
}
