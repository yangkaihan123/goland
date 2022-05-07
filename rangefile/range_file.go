package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := "C:\\Users\\ThinkPad\\Desktop\\1.txt"
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	reader := bufio.NewReader(file)
	for {
		buf := make([]byte, 1024)
		info, err := reader.Read(buf)
		fmt.Println(strconv.Itoa(info))
		fmt.Println(string(buf))
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read ok")
				break
			} else {
				fmt.Println("Read file error: ", err)
				return
			}
		}
	}
}
