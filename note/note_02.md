**Golang内置类型和函数**  
-

### 1.内置类型
值类型  

    `bool`  
    `int (32 or 64)`,`int8`,`int16`,`int32`,`int64`  
    `uint (32 or 64)`,`unit8 (byte)`,`unit16`,`unit32`,`unit64`  
    `float32`,`float64`  
    `string`
    `complex64`,`copmlex128`  
    `array`  -- 固定长度的数组

引用类型 （指针类型）  

    `slice` -- 序列数组（最常用）  
    `map`  -- 映射  
    `chan` -- 管道  

内置函数  

    `append`  -- 用来追加元素到数组，slice中，返回修改后的数组，slice  
    `close`  -- 主要用来关闭channel  
    `delete`  -- 从map中删除key对应的value  
    `panic`  -- 停止常规的goroutine (panic和recover:用来做错误处理)  
    `recover`  -- 允许程序定义goroutine的panic动作  
    `real`  -- 返回complex的实部 （complex, real imag:用于创建和操作复数）  
    `imag`  -- 返回complex的虚部  
    `make`  -- 用来分配内存，返回type本身（只能应用于slice,map,channel）
    `new`   -- 用来分配内存，主要用来分配值类型，比如 int,struct 返回指向type的指针  
    `cap`   -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片slice和map)  
    `copy`  -- 用于复制和连接slice,返回复制的数目  
    `len`   -- 来求长度，比如 string,array,slice,map,channel 返回长度  
    `print,println`  -- 底层打印函数，一般使用fmt包来调用  

### 内置接口 error
```go
type error interface {

        Error()  string

}
```
