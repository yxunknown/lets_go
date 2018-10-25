package main

import (
	"fmt"
	"math"
)

func main() {
	x := float64(-64)
	if sqrtx, err := sqrts(x); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt is", sqrtx)
	}

}

type ErrNegativeSqrt float64

func (fnum ErrNegativeSqrt) Error() string {
	//fmt包在输出时，除了会匹配String()方法外，也会尝试匹配Error()方法
	//对于值类型的接收者
	//fmt.Printf(fnum)等同于fmt.Printf(fnum.Error)
	//从而导致无限递归，死循环
	//use convert to avoid stack overflow err
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(fnum))
}

func sqrts(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}
