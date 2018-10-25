package main

/**
实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流
*/
import "fmt"

func main() {
	reader := MyReader{}
	buf := make([]byte, 8)
	for {
		n, err := reader.Read(buf)
		fmt.Printf("n = %v, err = %v, buf = %v\n", n, err, buf)
		if err == nil {
			fmt.Printf("buf[:n] is %q\n", buf[:n])
		} else {
			fmt.Println(err)
		}
	}
}

type MyReader struct{}

func (reader MyReader) Read(b []byte) (n int, err error) {
	//每次生成一个A，并放入buffer当中返回出去
	b[0] = byte('A')
	return 1, nil
}
