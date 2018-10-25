package main

import "fmt"

/**
发送者可以通过close来关闭一个channel，以表示没有数据继续输入到该channel

接收者可以利用
v, ok := <- ch
来判断channel是否被关闭
*/
func main() {
	ch := make(chan int, 10)
	go fibonacciChannel(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func fibonacciChannel(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
