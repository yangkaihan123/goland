**Map声明，元素访问以及遍历**
map的声明：  
```go
m := map[string]int{"one":1,"two":2,"three":3}
m1 := map[string]int{}
m1["one"]=1
m2 := make(map[string]int,10 /*Initial Capacity*/) //参数：map类型和cap
//为什么不初始化len
```   
map和slice都是可自增的  
go语言中的map在访问key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在  
```go
if v, ok := m["four"]; ok {
	t.Log("four",v)
}else{
	t.Log("not existing")
}
```   
map的遍历   
```go
m := map[string]int{"one":1,"two":2,"three":3}
for k,v := range m {
	t.Log(k,v)
}
```  