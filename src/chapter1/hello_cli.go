package main

import (
	"flag"
	"fmt"
)

var (
	name    = flag.String("name", "mevur", "the name will be greeted")
	spanish = flag.Bool("spanish", false, "whether use spanish")
	help    = flag.Bool("help", false, "help")
)

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	if *spanish {
		fmt.Println("Holo", *name)
	} else {
		fmt.Println("Hello", *name)
	}
}
