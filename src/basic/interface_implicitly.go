package main

/**
类型通过直接实现某些方法来实现特定的接口，无需在实现上添加接口的名称
隐式接口解耦了实现接口的包和定义接口的包，互不依赖
因此没有 implements 关键字
*/
import (
	"fmt"
	"os"
)

func main() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello word")
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}
