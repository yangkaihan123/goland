package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
	/*
		函数 flag.StringVar 接受四个参数
		第一个参数是用于存储该命令参数的地址，具体到这里就是前面声明的变量name的地址（内存里的位置），由表达式&name表示
		第二个参数是为了指定该命令参数的名称，这里是Name
		第三个参数是为了指定在未追加该命令的参数是的默认值，这里是everyone
		第四个参数是该命令参数的简短说明，这行在打印命令说明时会用到
		还有一个与flag.StringVar 类似的函数 flag.String 这两个函数的区别是，flag.String 会直接返回一个已经分配好的用于存储命令参数值的地址。
		如果使用 flag.String 的话 那么 var name 赋值的时候就赋值成：
		var name = flag.String("name","everyone","The greeting object")
	*/

}

func main() {
	flag.Parse()
	/*
		flag.Parse() 函数用于真正解析命令参数，并把它们的值赋给相应的变量
		对这个函数的调用必须在所有命令参数存储载体里声明，这里是对变量name的声明和设置（flag.StringVar）之后，并且在读取任何命令参数之前进行。
		所以把 flag.Parse() 放在main的第一行
		设置了规则以后 但是并没有真正的parse 只有真正的parse了 才能吧解析后的参数值装到对应的变量中
		所以给命令参数赋值的操作放在了init里
		先定义和赋值变量，再放入flag.Parse()
	*/
	fmt.Printf("Hello, %s!\n", name)
}
