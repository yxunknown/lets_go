package main

import (
	"fmt"
	"math/cmplx"
)

/**
basic types of Go
bool
string
int int8 int16 int32 int64 //有符号整数
uint uint8 uint16 uint32 uint54 //无符号整数
byte = uint8
rune = int32 //表示Unicode编码
float32 float64 //浮点数
complex64 complex128 //复数
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
