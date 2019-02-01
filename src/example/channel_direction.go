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
when using channels as parameters of function,
you can specify the channel can be used on receive only or send only,
or both.
*/
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "some-data")
	pong(pings, pongs)
	data := <-pongs
	fmt.Println(data)
}

// can use the channel send data only,
// if try to receive data from channel,
func ping(ping chan<- string, msg string) {
	// send data only on this channel
	ping <- msg + " after-ping"
}

// ping channel can receive data only
// pong channel can send data only
func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	msg += " after-ping-rev"
	pong <- msg + " after-pong-snd"
}
