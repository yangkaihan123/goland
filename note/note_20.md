**Go的相关接口**  
---
[interface_test.go](study_source/study_10/interface/interface_test.go)
接口和依赖  
java的接口依赖示意图: [java的接口与依赖.png](pic/note_20/java的接口与依赖.png)   
go的interface：   
接口定义   
```go
type Programmer interface {
	WriteHelloWorld() Code
}
```    
接口实现：  
```go
type GoProgrammer struct {
}
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}
```    
WriteHelloWorld() 这个方法绑定在了结构体上，接口实现里就是绑定在了GoProgrammer这个结构的指针   

Go接口  
---
差异：   
1.接口为非入侵性，实现不依赖于接口定义   
2.所以接口的定义可以包含在接口使用者包内   

接口变量    
---
```go
var prog Coder = &GoProgrammer{}
```
Coder是一个接口  GoProgrammer是一个实现  prog是类型是接口的变量  
当prog被初始化后，有两部分：类型和数据  
type GoProgrammer struct{
} 是实现接口的类型   
&GoProgrammer 是实现GoProgrammer的一个实例

自定义类型
---
[customer_type_test.go](study_source/study_10/customer_type/customer_type_test.go)   
1.type IntConvertionFn func(n int) int   
2.type MyPoint int