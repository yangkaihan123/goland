package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	//create listener
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // break programming
	}
	// listen from agent link
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // break programming
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // break programming
		}
		fmt.Printf("Received data %v", string(buf[:len]))
	}
}
