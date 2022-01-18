使用闭包调试
-
    在分析和调试复杂程序时，无数个函数在不同的代码文件中相互调用，如果这时候能够准确的知道哪个文件中的具体哪个函数在执行，对于调试就很有帮助
    可以使用runtime或者log包中的一些函数来实现这样的功能
    runtime 包中的Caller() 函数提供了相应的信息，因此可以在需要的时候实现一个where() 闭包函数来打印函数执行位置：
    ex:
        where := func() {
            _, file, line, _ := runtime.Caller(1)
            log.Printf("%s:%d",file,line)
        }
        where()
        //some code
        where()
        //some more code 
        where()
        
    也可以设置log包中的flag参数实现：
    log.SetFlags(log.Llongfile)
    log.Print("")
    或者使用一个更加简单的where函数：
    var where = log.Print
    func func1() {
        where()
        ... some code
        where()
        ... some code
        where()
    }