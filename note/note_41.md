切片
-
    切片（slice）是对数组一个连续片段的引用（被引用的数组被称为相关数组，通常是匿名的），所以切片是一个引用类型
    （更类似c/c++中的数组类型，或者python中的list类型）
    这个片段可以是整个数组，或者是由起始和终止的索引标识的一些项的子集
    需要注意的是，终止索引标识的项不包括在切片内，切片提供了一个相关数组的动态窗口

    切片是可以索引的，并且可以由len()函数获取长度
    
    给定项的切片索引可能比相关数组的相同元素的索引小，和数组不同的是，切片的长度可以在运行时修改，
    最小为0，最大为相关数组的长度：切片是一个长度可变的数组
    
    切片提供了计算容量的函数 cap() 可以测量切片最长可以达到多少：
    它等于切片从第一个元素开始，到相关数组末尾的元素个数。
    如果 s 是一个切片，cap(s) 就是从 s[0] 到数组末尾的数组长度。
    切片的长度永远不会超过它的容量，所以对于切片s 来说 0 <= len(s) <= cap(s) 这个不等式永远成立
    
    多个切片如果表示同一个数组的片段，它们可以共享数据，因此一个切片和相关数组的其他切片是共享存储的，
    相反，不同的数组总是代表不同的存储。
    数组实际上是切片的构建块。

    优点：因为切片是引用，所以不需要使用额外的内存，并且比使用数组更有效率
    声明切片的格式：
    var identifier []type （不需要说明长度）
    一个切片在未初始化之前的默认值为nil，长度为 0
    切片的初始化格式：
    var slice1 []type = arr1[start:end]
    这表示 slice1 是由数组arr1 从start索引到end-1索引之间的元素构成的子集（切分数组，start:end 被称为slice表达式）
    所以slice[0] 就等于 arr1[start]。这可以在arr1被填充前就定义好
    如果写成 var slice1 []type = arr1[:] 那么slice1 就等于完整的arr1数组（所以这种表示方式是
    arr1[0:len(arr1)] 的一种缩写）。另一种表述方式是：slice1 = &arr1
    arr1[2:] 和 arr1[2:len(arr1)] 相同，都包含了数组从第三个到最后的所有元素
    arr1[:3] 和 arr1[0:3] 相同，包含了从第一个到第三个元素（不包括第四个/不包含下标为三的元素）
    
    如果要去掉slice1的最后一个元素，只要slice1 = slice1[:len(slice1)-1]
    一个由数字1，2，3 组成的切片可以用 s := [3]int{1,2,3}[:] 生成
    简单的写法：s := []int{1,2,3}
    s2 := s[:] 是用切片组成的切片，拥有相同元素，但是两个切片仍然指向相同的相关数组
    一个切片 s 可以使用 s = s[:cap(s)] 来扩展切片的大小上限，如果再扩大的话就会导致运行时错误
    
    对于一个切片（包括string）
    s == s[:i] + s[i:] // i 是一个整数且 0 <= i <= len(s)
    len(s) <= cap(s) 
    这个状态总是成立的

    切片也可以用类似数组的方式初始化：
    var x = []int{2,3,5,7,11}
    这样就可以创建一个长度为5的数组，并且创建了一个相关切片

    切片在内存中的组织方式实际上是一个有3个域的结构体：
    指向相关数组的指针，切片长度以及切片的容量
    长度为2 容量为4的切片y:
    y[0] = 3 且 y[1] = 5
    切片y[0:4] 由元素 3，5，7，11组成
如图：[slice_in_mem.png](pic/note_41/slice_in_mem.png)  
例程 [array_slices_test.go](study_source/slice/array_slices_test.go)
        
    output:
    Slice at 0 is 2  
    Slice at 1 is 3  
    Slice at 2 is 4  
    The length of arr1 is 6  
    The length of slice1 is 3  
    The capacity of slice1 is 4  
    Slice at 0 is 2  
    Slice at 1 is 3  
    Slice at 2 is 4  
    Slice at 3 is 5  
    The length of slice1 is 4  
    The capacity of slice1 is 4  
    如果 s2 是一个 slice，你可以将 s2 向后移动一位 s2 = s2[1:]，但是末尾没有移动。切片只能向后移动，s2 = s2[-1:] 会导致编译错误。切片不能被重新分片以获取数组的前一个元素。
    注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针！！

    关于切片的容量和长度：
    切片的长度就是它所包含的元素个数。
    切片的容量是从它的第一个元素开始数，到其相关数组元素末尾的个数。

###将切片传递给函数
如果一个函数需要对数组做操作，最通常的做法是把参数声明为切片，当调用该函数时，把数组分片，创建为一个切片引用并传递给该函数：          
计算数组元素和的方法：  
```go
package main
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
    }
	return s
}
func main()  {
    var arr = [5]int{0,1,2,3,4}
	sum(arr[:])
}
```   
###用make()创建一个切片
当相关数组还没有定义时，可以使用make()函数来创建一个切片，同时也创建好了相关数组：  
    
    var slice1 []type = make([]type, len)  
也可以简写为 slice1 := make([]type, len) ,这里len是数组的长度也是slice的初始长度  
所以定义 s2 := make([]int 10) 那么 cap(s2) == len(s2) == 10  
make 函数接受两个参数：元素的类型以及切片元素的个数  
如果创建一个slice1，他不占用整个数组，而是只占以len 为个数个项，那么只要：   
slice1 := make([]type, int, cap)   
make 的使用方式是：func make([]T, len, cap) 其中cap是可选参数   
所以下面两种方式都可以生成切片：      
make([]int, 50, 100)  
new([100]int)[0:50]  
下图是make方法生成的切片的内存结构：![](pic/note_41/make_slice.png?raw=true)    
例程：[make_slice_test.go](study_source/slice/make_slice_test.go)
```go
package slice

import "testing"

func TestMakeSlice(t *testing.T) {
	var slice1 []int = make([]int, 10)
	//load the array/slice
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	//print the slice:
	for i := 0; i < len(slice1); i++ {
		t.Logf("Slice at %d is %d\n", i, slice1[i])
	}
	t.Logf("\nThe length of slice1 is %d\n", len(slice1))
	t.Logf("The capacity of slice1 is %d\n", cap(slice1))
}
```  
输出：

    === RUN   TestMakeSlice
    make_slice_test.go:13: Slice at 0 is 0
    make_slice_test.go:13: Slice at 1 is 5
    make_slice_test.go:13: Slice at 2 is 10
    make_slice_test.go:13: Slice at 3 is 15
    make_slice_test.go:13: Slice at 4 is 20
    make_slice_test.go:13: Slice at 5 is 25
    make_slice_test.go:13: Slice at 6 is 30
    make_slice_test.go:13: Slice at 7 is 35
    make_slice_test.go:13: Slice at 8 is 40
    make_slice_test.go:13: Slice at 9 is 45
    make_slice_test.go:15:
    The length of slice1 is 10
    make_slice_test.go:16: The capacity of slice1 is 10
    --- PASS: TestMakeSlice (0.00s)
    PASS

因为字符串是纯粹不可变的字节数组，它们也可以被切分为切片  

练习：主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片  
[fibonacci_funcarray_test.go](study_source/slice/fibonacci_funcarray_test.go)  

###new()和make()的区别  

两者都是在堆上分配内存，但是它们行为不同，适用于不同的类型   

- new(T) 为每一个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 \*T 的内存地址： 这种方法 **返回一个指向类型T，值为 0 的地址的指针**，它适用于值类型：如数组和结构体，相当于`&T{}`。
- make(T) **返回一个类型为 T 的初始值**, 只适用于三种内建的引用类型：slice , map , channel   

new函数分配内存，make函数初始化 下图给出了区别：  
![](pic/note_41/make_new.png?raw=true)  

在上图中：  
第一个图中   
```go
var p *[]int = new([]int) // *p == nil; with len and cap 0
p := new([]int)
```  

在第二个图中， `p := make([]int, 0)`，切片 已经被初始化，但是指向一个空数组。   
这两种方式实用性不高，下面的方法：  
```go
var v []int = make([]int, 10, 50)
```  
或者  
```go
v := make([]int, 10, 50)
```   
这样分配有一个50个 int 值的数组，并且创建了一个长度为10，容量为 50 的切片v，该切片指向数组的前10个元素。  
**练习** ： 给定 s := make([]byte, 5)，len (s) 和 cap (s) 分别是多少？s = s[2:4]，len (s) 和 cap (s) 又分别是多少？  
例程：[make_slice02_test.go](study_source/slice/make_slice02_test.go)   
**练习2** 假设 s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := s1[2:]，s2 的值是多少？如果我们执行 s2[1] = 't'，s1 和 s2 现在的值又分别是多少？   
例程：[practice_slice_test.go](study_source/slice/practice_slice_test.go)  

###多维切片

和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。   
通过分片的分片（或者切片的数组），长度可以任意动态变化，所以Go语言的多维切片可以任意切分。而且，内层的切片必须单独分配（通过make)。    

###bytes包

类型 `[]byte` 的切片很常见，Go中有一个 bytes 包专门用来解决这种类型的操作。    
包bytes和strings包很类似，而且它还包含了一个很有用的类型 Buffer:
```go
import "bytes"

type Buffer struct{
	...
}
```   
这是一个长度可以变的 bytes 的 buffer , 提供 Read 和 Write 方法，因为读写长度未知的 bytes 最好用 buffer。  

Buffer 也可以这样定义：`var buffer bytes.Buffer`。  
或者使用 new 获得一个指针：`var r *bytes.Buffer = new(bytes.Buffer)`    
或者通过函数：`func NewBuffer(buf []byte) *Buffer`, 创建一个Buffer对象并且使用 buf 初始化好，NewBuffer 最好用在从 buf 读取的时候使用。   

**通过buffer串联字符串**  
类似Java的StringBuilder类。   
例程中，创建一个buffer，通过 `buffer.WriteString(s)` 方法将字符串 s 追加到后面，最后再通过 `buffer.String()` 方法转换成 string：   
```go
var buffer bytes.Buffer
for {
	if s, ok := getNextString(); ok {
		//method getNextString() not shown here
		buffer.WriteString(s)
    }else{
	    break	
    }
}
fmt.Print(buffer.String(),"\n")
```   
这种实现方式比使用 `+=` 要更节省内存和CPU，尤其要串联字符串的数目特别多的时候。  

**练习**：给定切片 sl，将一个 []byte 数组追加到 sl 后面。写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。   
例程：[](study_source/slice_02/bytes_01_test.go)   
**练习**：把一个缓存 buf 分片成两个切片：第一个是前 n 个bytes，后一个是剩余的    

##for-range结构

这种构建方法可以用于数组和切片   
```go
for ix, value := range slice1 {
	...
}
```   
第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值，都是仅在 for 循环内部可见的局部变量。  
value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。   
例程：[slice_forrange_test.go](study_source/slice_02/slice_forrange_test.go)    
例程2：[slice_forrange2_test.go](study_source/slice_02/slice_forrange2_test.go)   
slice_forrange2_test.go 给出了一个关于字符串的例子，`_` 可以用于忽略索引。   
如果只需要索引，可以忽略第二个变量：  
```go
for ix := range seasons {
	fmt.Printf("%d",ix)
}
//output: 0 1 2 3
```    
如果需要修改 `seasons[ix]`的值可以使用这个版本。    

**多维切片下的 for-range:**  

通过计算行数和矩阵值可以方便的写出：   
```go
for row := range screen {
	for column := range screen[row] {
		screen[row][column] = 1
    }   
}
```    
**问题：** 假设现在如下数组：`items := [...]int{10, 20, 30, 40, 50}`    
a) 如果写了下面这个for，那么执行for循环后的`items`的值是多少：   
```go
for _, item := range items {
	item *= 2
}
```   
b) 如果a)无法正常工作，写一个for循环让值可以double    
例程：[for_range_test.go](study_source/slice_02/for_range_test.go)   

**练习**  :
a) 写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。

如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？

b) 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。   
程序：[sum_array_test.go](study_source/slice_02/sum_array_test.go)    

**练习二**： 写一个minSlice方法，传入一个int的切片并且返回最小值，再写一个maxSlice方法返回最大值   
程序：[min_max_test.go](study_source/slice_02/min_max_test.go)








