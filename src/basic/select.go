package main

import "fmt"

/**
select可以让一个goroutine在多个channel上工作
select会阻塞，当条件分支某一个分支可以执行时，则执行该分支；
当有多个分支可以执行时，则随机选择一个执行
*/
func main() {
	c := make(chan int)
	quit := make(chan int)
	//goroutine调用匿名函数
	go func() {
		for i := 0; i < 10; i++ {
			//从c channel获取值
			fmt.Println(<-c)
		}
		//0输入到quit channel
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	//while loop
	//until quit channel receive data, this loop will break
	//otherwise this loop will calculate fibonacci array continuously
	for {
		select {
		case c <- x: //x 输入到c channel
			x, y = y, x+y
		case <-quit: // quit channel接收到输入
			fmt.Println("quit")
			return
		}
	}
}
