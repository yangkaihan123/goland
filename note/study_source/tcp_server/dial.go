package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9002")
	// tcp ipv4
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "localhost:9002")
	checkConnection(conn, err)
	// udp
	//conn, err = net.Dial("tcp", "[fe80::480a:7413:86a6:bb84%7]:9002")
	//ipv6
	//checkConnection(conn,err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connections is made with %v\n", conn)
}
