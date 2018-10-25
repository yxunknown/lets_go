package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "The term anno Domini is Medieval Latin and means in the year of the Lord"
	count := wordCount(s)
	fmt.Println(count)
}

func wordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	counter := make(map[string]int)
	for _, word := range words {
		counter[strings.ToLower(word)] += 1
	}
	return counter
}
