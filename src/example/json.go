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
	"strconv"
	"strings"
)

func main() {
	var obj map[string]IPS
	var ipt = `
	{
  		"ip": "192.168.137.1"
	}
	`
	if err := json.Unmarshal([]byte(ipt), &obj); err != nil {
		panic(err)
	}
	fmt.Println(obj)
	for k, v := range obj {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println()
}

type IPS struct {
	A uint8
	B uint8
	C uint8
	D uint8
}

// define ip error
type IPError struct {
	msg string
}

// implement Error interface
func (ipErr IPError) Error() string {
	return ipErr.msg
}

// parse a ip str to IP struct
// ipStr must be formatted like 122.132.141.124,
// when format of ipStr is invalid, an IPError occurs.
func strToIp(ipStr string) (ip IPS, err error) {
	seg := strings.Split(ipStr, ".")
	// check whether str is a valid ip
	if len(seg) != 4 {
		return IPS{}, IPError{"Invalid ip"}
	}
	var newIp IPS
	if a, err := strconv.ParseUint(seg[0], 10, 8); err != nil {
		return IPS{}, err
	} else {
		newIp.A = uint8(a)
	}
	if b, err := strconv.ParseUint(seg[1], 10, 8); err != nil {
		return IPS{}, err
	} else {
		newIp.B = uint8(b)
	}
	if c, err := strconv.ParseUint(seg[2], 10, 8); err != nil {
		return IPS{}, err
	} else {
		newIp.C = uint8(c)
	}
	if d, err := strconv.ParseUint(seg[3], 10, 8); err != nil {
		return IPS{}, err
	} else {
		newIp.D = uint8(d)
	}
	return newIp, nil
}

// implement Unmarshaler interface
func (ip *IPS) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Replace(str, "\"", "", -1)
	if newIp, err := strToIp(str); err != nil {
		return nil
	} else {
		*ip = newIp
	}
	return nil
}
