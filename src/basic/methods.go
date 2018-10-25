package main

import "fmt"

/**
Go没有类，但是可以在结构体上定义方法
*/
func main() {
	car := Car{speed: 90, vendor: "BMW"}
	car.info()
	car.run()

	f := MyFloat(23.000223)
	f.show()
}

type Car struct {
	speed  int
	vendor string
}

/**
define function on struct type
func     (t *T)   name(parameters...) T
keyword  receiver                     return type
*/
func (car *Car) run() {
	fmt.Println("i am running at", car.speed)
}

func (car *Car) info() {
	fmt.Println("Made by", car.vendor)
}

//可以对任意类型定义方法
type MyFloat float64

//需要用结构体包装一下
//接收者为指针类型
//指针可以避免方法调用时使用拷贝值，也可以在方法中通过指针修改值
func (f *MyFloat) show() {
	fmt.Println(*f)
}
