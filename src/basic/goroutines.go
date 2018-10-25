package main

import (
	"fmt"
	"time"
)

/**
goroutine 是go运行时提供的轻量级线程
go func(...)
*/
func main() {
	go say("Hello Goroutines")
	//sleep for waiting say func execution finish
	//time.Sleep(1 * time.Second)
	//fmt.Println("Main func is done")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
