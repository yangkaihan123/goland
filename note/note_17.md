**Go语言的函数**   
函数：一等公民   
主要差异：  
1.可以有多个返回值   
2.所有参数都是值传递：slice,map,channel 会有传递引用的错觉，其实都是值传递   
slice背后对应的其实是一个数组，切片本身是一个数据结构，在这个数据结构里，包含了指向后面数组的指针，即便在传值的情况下，这个结构被复制到函数里了，  
在通过背后的指针操作具体值的时候，还是在同一个内存空间，实际上是结构被复制了，但是结构里包含的指针还是同一个后端数组   
3.函数可以作为变量的值   
4.函数可以作为参数和返回值   