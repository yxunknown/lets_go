package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(20)
	fmt.Println("generate a random number ", rand.Int())
}
