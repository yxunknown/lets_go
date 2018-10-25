package main

import "fmt"

func main() {
	//buffered channel
	//the buffer size is 2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
