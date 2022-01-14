内置函数
-
    Go中有一些不需要进行导入第三方包就可以操作的内置函数。
    它们有时可以针对不同的类型进行操作例如：
    len cap append
    或者是用于系统级的操作，如：panic
    
    内置函数有：
    close 
    用于管道通信

    len和cap  
    len用于返回某个类型的长度或数量（字符串，数组，切片，map和管道）
    cap是容量的意思，用于返回某个类型的最大容量（只能用于切片和map)
    
    new和make
    new和make都是用于分配内存
    new 用于值类型和用户定义类型，如自定义结构
    make 用于内置引用类型（切片，map和管道）
    它们的用法就是像是函数，但是将类型作为参数：
    new(byte),make(byte)
    new(T) 分配类型T的零值并返回其地址，也就是指向类型T的指针
    也可以用于基本类型：v := new(int) 
    make(T) 返回类型T初始化后的值，因此make比new进行更多的工作
    需要注意的是new() 是一个函数，注意必须有括号
    
    copy和append
    用于复制和连接切片
    
    panic和recover
    两者均用于错误处理机制
        
    print和println
    底层打印函数，一般是使用了fmt包
    
    complex real 和 imag
    用于创建和操作复数