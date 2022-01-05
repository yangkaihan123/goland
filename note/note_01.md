**1.Go语言的主要特征**
---
****1.自动立即回收  
2.更丰富的内置类型  
3.函数的多返回值  
4.错误处理  
5.匿名函数和闭包  
6.类型和接口  
7.自带协程并发编程  
8.反射  
9.语言交互性****  

Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：

    1）首字母可以是任意的Unicode字符或者下划线（_）  
    2）剩余字符可以是Unicode字符，下划线，数字  
    3）字符长度不限  

Go只有25个关键字  

    `break` `default` `func` `interface` `select`
    `case` `defer` `go` `map` `struct`
    `chan` `else` `goto` `package` `switch`
    `const` `fallthrough` `if` `range` `type`
    `continue` `for` `import` `return` `var`

Go还有37个保留字  
Constants:    

    `true`  `false`  `iota`  `nil`

Types:    

    `int` `int8`  `int16`  `int32`  `int64`  
     `uint`  `uint8`  `uint16`  `uint32`  `uint64`  `uintptr`
     `float32`  `float64`  `complex128`  `complex64`
     `bool`  `byte`  `rune`  `string`  `error`

Functions:   

    `make`  `len`  `cap`  `new`  `append`  `copy`  `close`  `delete`
    `complex`  `real`  `imag`
     `panic`  `recover`

### Go的可见性：
    1）声明在函数内部，是函数的本地值，类似private  
    2）声明在函数外部，是对当前包可见（包内所有.go文件都可见）的全局值，类似protect
    3）声明在函数外部而且首字母大写是所有包可见的全局值，类似public

### Go语言的声明方式主要有四种
    var（声明变量） const（声明常量） type（声明类型） func（声明函数）  

_Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，用来说明该文件属于哪个包(package)，package声明下来就是import声明，再下来是类型，变量，常量，函数的声明_  


