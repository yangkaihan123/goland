切片重组
-

切片创建的时候一般比相关数组小：
```go
slice1 := make([]type, start_length,capacity)
```    
其中`start_length`作为切片初始长度而`capacity`作为相关数组的长度    
这么做的优势在于切片在达到容量上限以后可以扩容。   
改变切片长度的过程被称为切片重组reslicing，写法：`slice1 = slice1[0:end]`，其中end是新的末尾索引（即长度）。   
将切片扩展1位：   
```go
s1 = s1[0:len(s1)+1]
```    
切片也可以反复扩展直到占据整个相关数组   
例程：[reslicing_test.go](study_source/slice_03/reslicing_test.go)   
另一个例子：  
```go
var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
var a = ar[5:7]
//reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
```   
将a重新分片：   
```go
a = a[0:4] 
//reference of subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5 
```   
**问题 7.7**   

1) 如果 s 是一个切片，那么 s[n:n] 的长度是多少？

2) s[n:n+1] 的长度又是多少？

例程：[resliceing_02_test.go](study_source/slice_03/resliceing_02_test.go)    
output:   

    === RUN   TestReslicing2
    resliceing_02_test.go:7: the slice s length is 5, the cap is 5
    resliceing_02_test.go:9: the slice s length is 0, the cap is 3
    resliceing_02_test.go:11: the slice s length is 1, the cap is 1
    --- PASS: TestReslicing2 (0.00s)
    PASS