**iota**
-
iota 是go语言中的常量计数器，只能在常量的表达式中使用  
iota 在const 关键字出现时将被重置为0  
const 中每新增一行常量声明，iota 计数一次（iota 可以理解为const 语句块中的行索引）  
使用iota 可以简化定义，定义枚举的时候很有用：  
```go
const (
	n1 = iota // 0
	n2        // 1
	n3        // 2
	n4        // 3
)
```   

常见的iota用法
使用 _ 跳过某些值：  
```go
const (
    n1 =  iota // 0
	n2         // 1
	_         
	n4         // 3
)
```   
iota 声明中间插队：  
```go
const (
    n1 = iota // 0
    n2 = 100  // 100
    n3 = iota  // 2
    n4         // 3
)
const n5 = iota // 0
```     
定义数量级 （这里的 << 表示左移操作，1 << 10 表示将 1 的二进制表示向左移 10 位，也就是由 1 变成了 10000000000，
也就是十进制的 1024 同理 2 << 2 表示将 2 的二进制表示向左移 2 位，也就是由10 变成了 1000 ，转换为十进制就是 8 ）  
```go
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
```    
多个iota定义在一行：   
```go
const (
	a, b = iota + 1, iota + 2  // 1, 2
	c, d                       // 2, 3
	e, f                       // 3, 4
)
```  

