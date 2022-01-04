**可变参数和defer**   
```go
func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
    }
	return s
}
```   
可变参数   

defer函数   
```go
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources")
    }()
    t.Log("Started")
	panic("Fatal error")  // defer仍会执行 ，panic抛出异常后defer还是会执行
}
```   
函数延迟执行，defer后不会立即执行，会return后执行   
