defer和追踪
-
###defer
    关键字defer允许程序推迟到函数返回之前（或任意位置执行return之后）一刻才执行某个语句或函数，因为return可以包含一些操作，不是单纯的返回某个值，
    所以得等到返回后执行
    
    关键字defer的使用类似于java里的finally，一般用来释放某些已经分配的资源
    ex:
    func TestDefer(t *testing.T) {
	    function1()
    }
    func function1() {
        fmt.Printf("In func1 at the top\n")
        defer function2()
        fmt.Printf("in func1 at the bottom!\n")
    }
    func function2() {
        fmt.Printf("func2: Deferred until the end of the calling func1!")
    }

    使用defer的语句同样可以接受参数：
    func a() {
        i := 0
        defer fmt.Println(i)
        i++
        return
    }
    这个函数在执行defer语句时打印为：0
    
    当有多个defer行为被注册的时候，会以逆序执行（类似栈，后进先出）
    func f() {
        for i := 0; i < 5; i++ {
            defer fmt.Printf("%d\n",i)
        }
    }
    这个函数执行会输出： 4 3 2 1 0
    (试了下，确实加defer是倒序，不加defer是0 1 2 3 4）
    
    关键字defer允许进行一些函数执行完的收尾工作：
    1：关闭文件流
    ex：
        //open a file
        defer file.Close()
    2：解锁一个加锁的资源
    ex：
        mu.Lock()
        defer mu.Lock()
    3：打印最终报告
    ex：
        printHeader()
        defer printFooter()
    4：关闭数据库连接
    ex：
        //open a databas connection
        defer disconnectFromDB()
    
    例程模拟了这四种情况：
    func TestDefer4(t *testing.T) {
	    doDBOperations()
    }

    func doDBOperations() {
        connectToDB()
        fmt.Println("Defering the database disconnect.")
        defer disconnectFromDB() //function called here with defer
        fmt.Println("Doing some DB operations ...")
        fmt.Println("Oops! some crash or network error ...")
        fmt.Println("Returning form function here!")
        return
        //terminate the program
        //deferred function executed here just before actually returning, even if
        //there is a return or abnormal termination before
    }
    func disconnectFromDB() {
        fmt.Println("ok, the database is disconnected")
    }
    func connectToDB() {
        fmt.Println("ok, the database is connected")
    }
###defer实现代码追踪
    一个基础但是十分实用的实现代码执行追踪的方案，就是在进入和离开某个函数打印相关信息，即可以提炼为下面两个函数：
    func trace(s string) {
        fmt.Println("entering:",s)
    }
    func untrace(s string) {
        fmt.Println("leaving:",s)
    }
    例程展示了合适的调用这两个函数：
    func trace(s string) {
	    fmt.Println("entering:", s)
    }
    func untrace(s string) {
        fmt.Println("leaving:", s)
    }
    func a() {
        trace("a")
        defer untrace("a")
        fmt.Println("in a")
    }
    func b() {
        trace("b")
        defer untrace("b")
        fmt.Println("in b")
        a()
    }
    func TestTrace(t *testing.T) {
        b()
    }
    例程的输出结果为：
    entering: b
    in b
    entering: a
    in a
    leaving: a
    leaving: b

    更简洁的写法：
    func trace_new(s string) string {
        fmt.Println("entering:", s)
        return s
    }
    func un(s string) {
        fmt.Println("leaving:", s)
    }
    func a() {
        defer un(trace_new("a"))
        fmt.Println("in a")
    }
    func b() {
        defer un(trace_new("b"))
        fmt.Println("in b")
        a()
    }
    func TestDefertrace(t *testing.T) {
        b()
    }

###使用defer来记录函数的参数和返回值
    例程展示了在调试时使用defer的手法：
    func func1(s string) (n int, err error) {
        defer func() {
            log.Printf("fun1(%q) = %d, %v",s,n,err)
        }()
        return 7,io.EOF
    }
    func TestDeferLogValues(t *testing.T)  {
        func1("Go")
    }
    例程输出为：2022/01/14 10:17:03 fun1("Go") = 7, EOF
