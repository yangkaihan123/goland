闭包 匿名函数
-
    当我们不希望给函数起名字的时候，可以使用匿名函数：
    ex: 
        func(x, y int) int {
            return x + y
        }
    这样一个函数不能独立存在，（报错:non-declaration statement outside function body)
    但是可以被赋值于某个变量，即保存函数的地址到变量中：
    fplus := func(x, y int) int { return x + y }
    然后通过变量名对函数进行调用：fplus(3,4)
    也可以直接调用匿名函数：func(x, y int) int { return x + y } (3,4)
    ex:
        func() {
            sum := 0
            for i := 1; i <= 1e6; i++ {
                sum += 1
            }
        }()
    计算从一到一百万整数的和的匿名函数
    这里参数列表的第一对括号必须紧靠关键字func,因为匿名函数没有名称,花括号{} 涵盖着函数体，最后一个括号表示对该匿名函数的调用
    
    ex:
        func TestLiteral(t *testing.T) {
	        f()
        }
        func f() {
            for i := 0; i < 4; i++ {
                g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
                g(i)
                fmt.Printf("- g is of type %T and has value %p\n", g, g)
            }
        } 
    例程中，g代表的是func(int)，变量的值是一个内存地址，所以实际上我们拥有的是一个函数值：匿名函数可以赋值给变量，并作为值使用
    
    练习：写一个用于打印 Hello World 字符串的匿名函数，并赋值给变量fv，然后调用该函数并打印fv的类型
    func TestFv(t *testing.T) {
	    fv := func() {
		    fmt.Println("hello world")
	    }
	    fv()
    	t.Logf("the type of fv is %T", fv)
    }
    输出为：hello world 
           the type of fv is func()

    匿名函数与普通函数一样可以接受或者不接受参数，
    ex: 
        func (u string) {
            fmt.Println(u)
            ...
        }(v)
    向匿名函数传参

###defer和匿名函数
    ex:
        func TestReturnDefer(t *testing.T) {
            t.Log(f2())
        }
        func f2() (ret int) {
            defer func() {
                ret++
            }()
            return 1
        }
    函数f返回时，ret的值为2，因为ret++ 是在执行return1 语句后发生的
    这个方式可以用于返回语句之后修改返回的error时使用

    关键字defer 经常配合匿名函数使用，可以用于改变函数的命名返回值
    匿名函数还可以配合go关键字来作为gorountine来用
    
    匿名函数同样被称为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。
    闭包可使得某个函数捕捉到一些外部状态，比如：函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声时的作用域。
    这种状态（作用域内的变量）都被共享到闭包环境中，因此这些变量可以在闭包内被操作，直到被销毁。
    闭包经常被用作包装函数，它们会预先定义好一个或多个参数用于包装。匿名函数另一个很好的应用就是使用闭包来完成更加简洁的错误检查。
