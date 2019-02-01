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
by default, channels are unbuffered,
sends(chan <-) are accepted when receivers are ready,
receives(<- chan) are accepted when sends are ready.

buffered channels have a buffer area
*/

func main() {
	// make a buffered channel
	// the 2nd param is buffer size
	ch := make(chan string, 2)
	// put data into channel, and wait to consume
	ch <- "the one"
	ch <- "the two"
	// channel buffer is full, will produce error
	//ch <- "the three"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// ch is empty, will produce error
	//fmt.Println(<- ch)
}
