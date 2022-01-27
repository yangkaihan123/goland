package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//open link
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		// 由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return //终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s", clientName)
	trimmedClient := strings.Trim(clientName, "\r\n")
	// windows 平台用 "\r\n" Linux 下用"\n"
	//给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--s%--", input)
		// fmt.Printf("trimmedInput:--s%--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + "says: " + trimmedInput))
	}
}
