package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	b := make([]byte, 8)
	//while loop
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
		/**
		n = 8, err = <nil>, b = [104 101 108 108 111 32 119 111]
		b[:n] = "hello wo"
		n = 3, err = <nil>, b = [114 108 100 108 111 32 119 111]
		b[:n] = "rld"
		n = 0, err = EOF, b = [114 108 100 108 111 32 119 111]
		b[:n] = ""
		*/

	}
}
