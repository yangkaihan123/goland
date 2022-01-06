**if-else**
-
    当if结构内有break,continue,goto,return语句时,go代码的常见写法是省略else部分：  
```go
if condition {
	return x
}
return y
```   
    需要注意的是不要同时在if-else结构的两个分支都使用return，会编译错误  

###关于if-else的实用例子
    1.判断一个字符串是否为空：
        。 if str == "" { ... }
        。 if len(str) == 0 { ... }  使用len()函数来判断字符串是否为空
    2.判断运行Go程序的操作系统类型，可以用runtime.GOOS来判断   
        if runtime.GOOS == "Windows" {
            ...
        } else {
            ... 
        }
        
    3.函数Abs()用于返回一个整型数字的绝对值
        func Abs(x int) int {
            if x < 0 {
                return -x
            }
            return x
        }
    
    4.isGreater() 可以用于比较两个整型数的大小
        func isGreater(x, y int) bool {
            if x > y {
                return true
            }
            return false
        }