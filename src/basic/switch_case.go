package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go is now running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
	whenIsSaturday()
	greeting()
}

func whenIsSaturday() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	//switch cases is ordered execution
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	default:
		fmt.Println("Too far away")
	}
}

func greeting() {
	t := time.Now().Hour()
	//没有条件的switch结构
	//用于实现多层if-else结构
	switch {
	case t < 12:
		fmt.Println("Good morning.")
	case t < 17:
		fmt.Println("Good afternoon.")
	case t < 22:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good night.")
	}
}
