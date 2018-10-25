package main

import "fmt"

/**
接口是由一组方法所定义的结合
*/
type Abser interface {
	abs() float64
}

type AbsFloat float64

//实现接口方法直接进行定义即可
func (f *AbsFloat) abs() float64 {
	if *f < 0 {
		return float64(-*f)
	} else {
		return float64(*f)
	}
}
func main() {
	var a Abser
	pf := AbsFloat(231.2314)
	nf := AbsFloat(-23.12314)
	fmt.Println(pf.abs())
	fmt.Println(nf.abs())
	//abs只在AbsFloat指针类型上实现，因此无法对普通类型进行引用
	//a = pf
	//fmt.Println(a.abs())
	a = &nf
	fmt.Println(a.abs())
}
