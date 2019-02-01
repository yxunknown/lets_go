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
	"time"
)

/**
select can make goroutines work on multiple channels at same time
*/

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "msg1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "msg2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	out := make(chan int)
	sync := make(chan bool)
	go func() {
		for i := 0; i <= 10; i++ {
			x := <-out
			fmt.Print(x, " ")
		}
		sync <- true
	}()
	fbc(out, sync)
	fmt.Println()
}

// synchronize between goroutines
func fbc(out chan int, sync chan bool) {
	x, y := 0, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-sync:
			return
		}
	}
}
