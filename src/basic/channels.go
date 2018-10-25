package main

import "fmt"

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	//make() 可以用来创建任何类型的变量
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/**
name chan T
define an channel
*/
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	/**
	像sum送入c
	类似于c++的流操作
	cin >> c  // 输入流送到变量
	cout << c // 变量送到输入流
	*/
	c <- sum
}
