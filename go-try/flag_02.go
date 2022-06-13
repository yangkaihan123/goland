package main

import (
	"flag"
	"fmt"
	"os"
)

var name02 string

func init() {
	flag.StringVar(&name02, "name", "everyone", "The greeting object")
}
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
		flag.PrintDefaults()
	}
	/*
			自定义命令源码文件的参数使用说明
			最简单的方式就是对变量 flag.Usage 重新赋值。
			flag.Usage 的类型是 func() 是一个无参数声明且无结果声明的函数类型
			flag.Usage 变量在声明时就已经被赋值了，所以才能在运行go run flag_02.go --help 的时候才会出现：
			Usage of question
		  		-name string
					The greeting object (default "everyone")

			注意：对flag.Usage 的赋值必须在调用 flag.Parse() 函数之前
	*/
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name02)
}
