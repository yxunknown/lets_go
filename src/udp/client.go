package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

var input string

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6666")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	for {
		fmt.Scanln(&input)
		if _, err := Conn.Write([]byte(input)); err != nil {
			println("error occurs when sending data:", err)
		}
	}
}
