package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":6666")
	handleErr(err)
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	handleErr(err)
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	fmt.Printf("start listening on 0.0.0.0:6666\n")
	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err == nil && n > 0 {
			// receive data
			go func(data string) {
				forward(data)
			}(string(buf[0:n]))
		}
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Something bad occurs", err)
		os.Exit(0)
	}
}

func forward(data string) {
	fmt.Println("Forwarding data", data)
}
