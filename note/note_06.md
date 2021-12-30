**下划线**  
"_" 下划线是特殊标识符，用来忽略结果  

1.下划线在"import" 中  
在golang里，import的作用是导入其他package。  
import 下划线（如：import hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数

示例：  
代码结构：  
```shell
    src 
    |
    +--- main.go            
    |
    +--- study_source
           |
            +--- study_source.go
```
```go
package main

import _"./study_source"

func main()  {
    //study_source.Print()
	//编译报错： ./main.go:6.5: undefined: study_source
}
```
hello.go  
```go
package hello

import "fmt"

func init()  {
    fmt.Println("imp-init() come here.")
}
func Print() {
    fmt.Println("study_source!")
}
```
输出结果：  
```shell
imp-init() come here.
```

2.下划线在代码中  
```go
package main

import (
	"os"
)

func mian() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/Balin/Desktop/text.txt")
	defer f.Close()
	for  {
        n, _ := f.Read(buf)
		if n == 0 {
			break
			
			
        }
		os.Stdout.Write(buf[:n])
	}
}
```
代码解析1：  
下划线的意思是忽略这个变量  
比如 os.open 返回值为 *os.File, error  
普通写法是 f, err := os.Open("xxxxx")  
如果此时不需要知道返回的错误值  
就可以用 f, _ := os.Open("xxxxx")  
如此便忽略了error变量  

代码解析2：  
占位符，意思是那个位置本应该赋值给某个值，但是不需要这个值  
就把该值赋值给下划线，意思是丢掉不要  
这样编译器更好的优化，任何类型的单个赋值都可以丢给下划线  
这种情况是占位用的，方放返回两个结果，但是只想要一个结果，就可以把不想要的那个值丢给下划线  

下划线的第三种用法：  
```go
import (
    "database/sql"
	_"github.com/go-sql-driver/mysql"
)
```  
第二个import就是不直接使用mysql包，只是执行一下这个包的init函数，把mysql的驱动注册到sql包里，然后程序里就可以使用sql包来访问数据库了  
