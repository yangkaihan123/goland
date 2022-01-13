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
    