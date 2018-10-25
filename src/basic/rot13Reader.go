package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot := Rot13Reader{s}
	io.Copy(os.Stdout, &rot)
}

type Rot13Reader struct {
	r io.Reader
}

func (rot *Rot13Reader) Read(buf []byte) (n int, err error) {
	/**
	rot13对换原理
	对于A-M和a-m的字符，ascii编码 加13 即可
	对于N-Z和n-z的字符，ascii编码 减13 即可
	对于非字母字符不做处理
	*/
	//从rot内部的reader读取流
	n, err = rot.r.Read(buf)
	//正常取到内容
	if err == nil {
		//对读取到的内容进行处理
		//只处理有效数据（0-n个字符）
		for index := 0; index < n; index++ {
			value := buf[index]
			if (value >= byte('A') && value <= byte('M')) ||
				(value >= byte('a') && value <= byte('m')) {
				//A-M和a-m的字符，ascii编码 加13
				buf[index] = value + 13
			} else if (value >= byte('N') && value <= byte('Z')) ||
				(value >= byte('n') && value <= byte('z')) {
				//N-Z和n-z的字符，ascii编码 减13
				buf[index] = value - 13
			}
		}
	}
	return
}
