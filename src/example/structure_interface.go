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
	"fmt"
	"strconv"
)

func main() {
	ip := IP{192, 168, 138, 64}
	toString(23)
	fmt.Println(toString(132))
	fmt.Println(toString(ip))
}

type Stringer interface {
	string() string
}

func toString(any interface{}) string {
	// try to convert any to Stringer
	// if convert success, v is Stringer type and ok is true
	// visa verse, v is nil and ok is false
	if v, ok := any.(Stringer); ok {
		return v.string()
	}
	switch v := any.(type) {
	case int:
		// format int64
		return strconv.Itoa(v)
	default:
		return "???"
	}
}

type IP struct {
	A uint8
	B uint8
	C uint8
	D uint8
}

func (ip IP) string() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip.A, ip.B, ip.C, ip.D)
}
