package main

import "fmt"

func main() {
	ip := IPAddr{10, 1, 124, 111}
	fmt.Println(ip)
	addrs := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"GoogleDNS": {8, 8, 8, 8},
	}
	fmt.Println(addrs)
	for key, value := range addrs {
		fmt.Println(key, value)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := "%v.%v.%v.%v"
	return fmt.Sprintf(s, ip[0], ip[1], ip[2], ip[3])
}
