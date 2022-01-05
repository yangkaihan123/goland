**面向对象**
-
###1.行为的定义和实现   
官方对go的面向对象的说法：   

    Is Go an object-oriented language?   
    Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).
    Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.   

    Go 是一种面向对象的语言吗？也是也不是。尽管 Go 有类型和方法，并允许面向对象的编程风格，但没有类型层次结构。 Go 中“接口”的概念提供了一种不同的方法，我们认为它易于使用，并且在某些方面更通用。
    还有一些方法可以将类型嵌入到其他类型中，以提供与子类化类似但不完全相同的东西。此外，Go 中的方法比 C++ 或 Java 中的方法更通用：它们可以为任何类型的数据定义，甚至是内置类型，例如普通的“未装箱”整数。它们不限于结构（类）。此外，由于缺乏类型层次结构，Go 中的“对象”感觉比 C++ 或 Java 等语言轻得多。
    没有继承，不支持继承  

###**封装数据和行为**
结构体定义：  
```go
type Employee struct{
    Id    string
	Name  string
	Age   int
}
```   
实例创建和初始化：   
```go
e := Employee{"0","Bob","20"}
e1 := Employee{Name: "Mike",Age: 30}
e2 := new(Employee) // new创建了一个指向实例的指针，这里返回的引用或者指针，相当于e := &Employee{}
e2.Id = "2"
e2.Age = 22
e2.Name = "Rose"
```    
###**行为（方法）的定义**    
```go
type Employee struct{
    Id    string
	Name  string
	Age   int
}
```    
```go
// 第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.ID,e.Name,e.Age)
}
```   
```go
//通常情况下为了避免内存拷贝使用下面方式定义
func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}
```    
两种方法的写法，第一种是复制的方式，第二种是直接操作指针   