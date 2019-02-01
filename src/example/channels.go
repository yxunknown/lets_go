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

/**
channels are pipes the connect different concurrent goroutines.
*/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// use make(chan value-type) to create new channel
	// <- ch get data from channel
	// ch <- input data
	// in general, send and receive operation will block
	// until both receiver and sender are ready. this property
	// can be used to achieve synchronize between goroutines.
	ch := make(chan string)
	go useChannel(ch)

	for i := 0; i < 5; i++ {
		ch <- "give you a " + strconv.Itoa(i)
	}
	ch <- "q"
}

func useChannel(ch chan string) {
	for {
		msg := <-ch
		if strings.ToUpper(msg) == "Q" {
			break
		}
		fmt.Println("receive msg:", msg)
	}
}
