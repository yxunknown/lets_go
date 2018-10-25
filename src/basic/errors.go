package main

import (
	"fmt"
	"time"
)

/**
error是一个内建的接口，用于表示错误状态
type error interface {
	Error() string
}
*/
/**
对于返回error的函数
应该对error进行处理
如果error为nil，则表示函数正常执行
否则表示函数执行发生错误
*/
type MyError struct {
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", err.When, err.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did'nt work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
