**变量和常量**
1.变量  
变量：  
    程序运行过程中的数据都是保存在内存中，想要在代码中操作某个数据时就需要去内存上找到这个变量，但是  
    直接在代码中通过内存地址去操作变量的话，代码的可读性会变差，而且还容易出错，所以使用变量将这个数据  
    的内存地址保存起来，直接通过这个变量就能找到内存上对应的数据了。  

变量类型：  
变量（variable）的功能是存储数据，不同的变量保存的数据类型可能就会不一样  
高级编程语言基本形成了一套固定的类型，常见变量的数据类型有：整型，浮点型，布尔型等。  
GO 语言中的每一个变量都有自己的类型，并且变量必须经过声明才可以开始使用。  

变量声明：  
Go语言中的变量需要声明后才可以使用，同一作用域内不支持重复声明，并且Go语言的变量声明后必须使用。  

标准声明：  
Go语言中的变量声明格式为：
```go
var 变量名 变量类型
```
变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号， ex:   
```go
var name string
var age int
var isOK bool
```  
批量声明变量   
每声明一个变量就需要写var关键字，比较麻烦，go中支持批量声明变量：  
```go
var (
	a string
	b int
	c bool
	d float32
)
```  
变量的初始化  
go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作，每个变量都会被初始化成其类型的默认值  
例如： 整型 int 和 浮点型 float 变量的默认值为0，字符串变量默认值为空字符串，布尔型变量的默认值为false  
切片（slice）函数（func） 指针变量的默认值为nil  
也可以在声明变量的时候直接指定初始化值：  
```go
var 变量名 类型 = 表达式
ex:  
var name string = "pprof.cn"  
var sex int = 1
//也可以一次初始化多个变量
var name, sex = "pprof.cn", 1
```  
类型推导：  
有时候会将变量的类型省略，这个时候，编译器会根据等号右边的值来推导变量的类型完成初始化  
```go
var name = "pprof.cn"
var sex = 1
```  
短变量声明：  
在函数内部，可以使用更加简单的方式声明并初始化变量。 ex:  
```go
package main

import (
	"fmt"
)

//全局变量
var m = 100

func main()  {
    n := 10
	m := 200 //此处声明局部变量
	fmt.Println(m,n)
}
```  
匿名变量：  
在使用多重赋值的时候，如果想要忽略某个值，可以使用匿名变量，匿名变量使用一个下划线表示：  
```go
func foo() (int, string) {
	return 10, "Q1mi"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=",x)
	fmt.Println("y=",y)
}
```   
匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不会存在重复声明   
注意事项：  
函数外的每个语句都必须以关键字开始（var, const, func等）  
:= 不能用在函数外  
_ 多用于占位，表示忽略值  