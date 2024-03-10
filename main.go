package main

import (
	"fmt"
	"net"	
	"os"
)

func main() {
	port := os.Args[1]
//	prefix := os.Args[2]

	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)	
	if err != nil {
		fmt.Println("Failed to start listener", err)
	}
	fmt.Printf("Started listening at %s", listener.Addr())

	listener.Accept()
}
