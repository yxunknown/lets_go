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
go support anonymous functions

闭包：
[函数]和[函数内部能访问的变量]的总和就是一个闭包。

闭包的作用：间接访问一个变量，也就是隐藏变量

函数返回函数的闭包方式：
为了避免某个变量变为全局变量，通常会将该变量放入一个函数里，再在该函数内部
定义另一个函数访问该变量，此时就会形成一个闭包。该闭包的作用就是为了「隐藏变量」

内存泄露的定义：某些变量无法再被使用（访问不到）却占有内存。

*/
func main() {
	nextInt := intSeq()

	go seq(nextInt, "1")
	go seq(nextInt, "2")
	go seq(nextInt, "3")

	time.Sleep(time.Second * 2)
}

func seq(f func() int, name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("name is", name, "value is", f())
	}

}

func intSeq() func() int {
	// i 变量和返回的匿名函数组成了一个闭包
	var i = 0
	return func() int {
		i++
		return i
	}
}
