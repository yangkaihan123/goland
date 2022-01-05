**循环**  
-
Go仅支持循环关键字for  
```go
//while 条件循环
//while (n < 5)
n := 0
for n < 5 {
	n++
	fmt.Println(n)
}
//while循环也使用关键字for
```  
```go
//无限循环 while (true)
n := 0
for {
	...
}
```   
###**条件语句**
```go
if condition {
	// code of be executed if condition is true 
} else {
	// code of be executed if condition is false
}
```   
```go
if condition-1 {
	// code to be executed if condition-1 is true
} else if condition-2 {
	// code to be executed if condition-2 is true
} else {
	// code to be executed if both  condition-1 and condition-2 are false
}
```   
主要的差异：  
1. condition 表达式结果必须是布尔值  
2. 支持变量赋值：
```go
if var declaration; condition {
	// code to be executed if condition is true
}
```  
switch条件的区别  
```go
switch os := runtime.GOOS; os {  // 不限制常量或者整数
case "darwin":
	fmt.Println("OS X.")
	//break  //不需要加break
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd,openbsd,
		// plan9,windows...
		fmt.Println("%s.",os)
}
```  
```go
switch {
    case 0 <= Num && Num <= 3:
		fmt.Printf("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Printf("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Printf("7-9")
}
```  
1.条件表达式不限制为常量或者整数  
2.单个case中，可以出现多个结果选项，使用逗号分隔   
3.Go语言不需要用break来明确退出一个case  
4.可以不设定switch之后的条件表达式，这种情况下整个switch结构与多个if...else...的逻辑作用等同  
