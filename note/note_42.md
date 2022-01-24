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

切片的复制与追加
-

如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容全部都拷贝过来，例程演示了从拷贝切片的 copy 函数，和向切片追加新元素的 append 函数   
例程： [copy_append_test.go](study_source/slice_03/copy_append_test.go)   
`func append(s[]T, x ...T) []T` 其中append函数将0个或者多个具有相同类型的 s 的元素追加到切片后面并且返回新的切片    
追加的元素必须和原切片的元素同类型，如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。   
因此，返回的切片可能已经指向一个不同的相关数组。append 总会执行成功，除非系统内存耗尽了   

如果想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可： `x = append(x,y...)`    
需要注意的时，append在大多数情况下都很好用，但是如果想完全掌控整个追加的过程，可以自己新建一个方法。   
例程：[append_test.go](study_source/slice_03/append_test.go)   
主要是AppendByte 方法    

`func copy(dst, src []T) int`  copy 方法将类型为T 的切片从源地址src拷贝到目标地址dst，覆盖dst的相关元素，并且返回拷贝元素的个数。   
源地址和目标地址可能会有重叠。拷贝的个数是src和dst的长度最小值。如果src是字符串嘛呢元素类型就是byte。    
如果还想继续使用src，在拷贝结束后执行 `src = dst`。   

**练习 7.9**   
给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。      

**练习 7.10**    
用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。    
[filter_slice_test.go](study_source/slice_03/filter_slice_test.go)     
**练习 7.11**
写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。      
[Insert_Slice_test.go](study_source/slice_03/Insert_Slice_test.go)    
**练习 7.12**   
写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。    
[remove_slice_test.go](study_source/slice_03/remove_slice_test.go)   































