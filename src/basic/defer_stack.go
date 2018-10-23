package main

import "fmt"

/**
defer will push the defer-statement into a stack in order
when outer func has been done
the stack will pop the defer-statement in LIFO(last in first out)
order and execute them
*/
func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}
