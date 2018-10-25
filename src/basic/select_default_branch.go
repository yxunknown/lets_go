package main

import (
	"fmt"
	"time"
)

/**
当select的所有分支都没有准备好的时候
select会转而执行default分支
*/
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	//while loop
	for {
		//like switch
		select {
		case <-tick: //从tick接收数据
			fmt.Println("tick.")
		case <-boom: //从boom接收到数据
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
