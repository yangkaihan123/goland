数组与切片
-
#说明
    数组和切片还有Map都是容器，包含了大量条目的数据结构
    
    以 [] 符号标识的数组类型

###数组的概念
    数组是具有相同唯一类型的一组以编号且长度固定的数据项序列（同构的数据结构）
    这种类型可以是任意的原始类型，比如整型 int ,字符串 string 或者自定义类型
    数组的长度必须是一个常量表达式，并且必须是一个非负整数
    数组的长度也是数组类型的一部分，所以[5]int 和 [10]int 是不同类型的
    
    数组编译时值的初始化时按照数组顺序完成的，如果想让数组元素类型为任意类型的话，可以给空接口作为类型。当使用值时必须先做一个类型判断

    数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素的索引为0，第二个索引为1
    元素的数目（长度）或者数组的大小必须是固定的并且声明在该数组时就给出，因为编译的时候需要知道数组的长度，以便分配内存，数组最大长度为2GB

    声明格式： var identifier [len]type
    ex: var arr1 [5]int
    数组arr1 在内存的结构为：
    arr1  [][][][][]
    index 0 1 2 3 4
    每个元素是一个整型值，当声明数组时，所有的元素都会被初始化为默认值0
    arr1 的长度为5，索引范围从 0 到 len(arr1)-1
    第一个元素 arr1[0], 第三个元素为 arr1[2], 索引i的元素是 arr1[i],最后一个元素为arr1[len(arr1)-1]
    对索引项为i的数组元素赋值：
    arr1[i] = value,所以数组是可变的
    只有有效的索引可以被使用，当使用等于或者大于len(arr1)的索引时如果编译器可以检测到，会给出索引超限的提示信息runtime error: index out of range；如果检测不到的话编译会通过而运行时会 panic
    
    由于索引的存在，遍历数组的方法是使用for结构
    。通过for 初始化数组项
    。通过for 打印数组元素
    。通过for 依次处理元素
    ex：
        func TestForArray(t *testing.T) {
	        var arr1 [5]int
	        for i := 0; i < len(arr1); i++ {
		        arr1[i] = i * 2
	        }
	        for i := 0; i < len(arr1); i++ {
		        t.Logf("Array1 at the index %d is %d\n", i, arr1[i])
	        }
        }
    for 循环中的条件非常重要：i < len(arr1) 如果写成 i <= len(arr1) 会产生越界错误

    遍历数组：
    for i:=0; i < len(arr1); i++ {
        arr1[i] = ...
    }
    也可以使用for-range 的方式：
    for i, _ := range arr1 {
        ..
    }
    这里i也是数组的索引，这两种方式也适用于切片（slice）
    
    练习1：下面代码输出是什么？
    a := [...]string{"a","b","c","d"}
    for i := range a {
        fmr.Println("Array item", i ,"is",a[i])
    }
    输出为：
    === RUN   TestArray2
    for_arrays_test.go:18: Array item 0 is a
    for_arrays_test.go:18: Array item 1 is b
    for_arrays_test.go:18: Array item 2 is c
    for_arrays_test.go:18: Array item 3 is d
    --- PASS: TestArray2 (0.00s)
    PASS
    
    Go语言中的数组是一种值类型，所以可以通过 new() 来创建：
    var arr1 = new([5]int) 
    这种方式创建的数组和 var arr2 [5]int 创建的数组，arr1 的类型是 *[5]int ,arr2 的类型是 [5]int
    这样的结果就是当把一个数组赋值给另一个时，需要在做一次数组内存的拷贝 
    ex:     
        arr2 := *arr1
        arr2[2] = 100
    这样两个数组是不同的值，在赋值后修改arr2不会对arr1生效，所以在函数中数组作为参数传入时，func1(arr2) 
    会产生一次数组拷贝，func1方法不会修改原始数组arr2
    如果想修改源数组，那么arr2必须通过 & 以引用的方式传过来，func1(&arr2)
    
    练习1：证明当数组赋值时，发生了数组内存拷贝  
    func TestArrayPointer(t *testing.T) {
	    var arr1 [5]int

	    for i := 0; i < len(arr1); i++ {
		    arr1[1] = i * 2
	    }
	    arr2 := arr1
	    arr2[2] = 100

	    for i := 0; i < len(arr2); i++ {
		    t.Logf("Array arr1 at index %d is %d\n", i, arr1[i])
	    }
	    t.Log()
	    for i := 0; i < len(arr2); i++ {
		    t.Logf("Array arr2 at index %d is %d\n", i, arr2[i])
	    }
    }
    输出结果：
    === RUN   TestArrayPointer
    array_value_test.go:15: Array arr1 at index 0 is 0
    array_value_test.go:15: Array arr1 at index 1 is 8
    array_value_test.go:15: Array arr1 at index 2 is 0
    array_value_test.go:15: Array arr1 at index 3 is 0
    array_value_test.go:15: Array arr1 at index 4 is 0
    array_value_test.go:17: 
    array_value_test.go:19: Array arr2 at index 0 is 0
    array_value_test.go:19: Array arr2 at index 1 is 8
    array_value_test.go:19: Array arr2 at index 2 is 100
    array_value_test.go:19: Array arr2 at index 3 is 0
    array_value_test.go:19: Array arr2 at index 4 is 0
    --- PASS: TestArrayPointer (0.00s)
    PASS
    证明了数组赋值，是修改了数组的指针，arr1的值未改

    练习2：写一个循环并用下标给数组赋值（从0到15）并且将数组打印在屏幕
    func TestForArray_02(t *testing.T) {
	    var a [15]int
	    for i := 0; i < len(a); i++ {
		    a[i] = i
	    }
	    t.Log(a)
    }
    输出：
    === RUN   TestForArray_02
    for_array_02_test.go:10: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
    --- PASS: TestForArray_02 (0.00s)
    PASS

    练习3：使用数组计算斐波那契数列，并打印前50个Fibonacci数字
    var fibs [50]int64
    func TestFibonacciArray(t *testing.T) {
	    fibs[0] = 1
	    fibs[1] = 1
	    for i := 2; i < 50; i++ {
		    fibs[i] = fibs[i-1] + fibs[i-2]
	    }
	    for i := 0; i < 50; i++ {
		    t.Logf("The %d-th Fibonacci number is %d\n", i, fibs[i])
	    }
    }
    这种方式很快，为什么