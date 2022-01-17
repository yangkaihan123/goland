应用闭包：将函数作为返回值
-
    先上例程：
    func TestFunctionReturn(t *testing.T) {
	// make an Add2 function, give it a name p2, and call it
	    p2 := Add2()
	    t.Logf("Call Add2 for 3 gives: %v\n", p2(3))
	//make a special Adder function, a gets value 2:
	    TwoAdder := Adder(2)
	    t.Logf("The result is %v\n", TwoAdder(3))
    }
    func Add2() func(b int) int {
        return func(b int) int {
            return b + 2
        }
    }
    func Adder(a int) func(b int) int {
        return func(b int) int {
            return a + b
        }
    }
    例程中，函数Add2 和 Adder 返回签名都是 func(b int) int 的函数：
    func Add2() (func(b int) int)
    func Adder(a int) (func(b int) int)
    Add2函数不接受任何参数，但是函数Adder接受一个int类型的整数作为参数，TwoAdder 存储了Adder 返回的参数
    
    略微不同的实现：
    func TestFunctionClosure(t *testing.T) {
	    var f = Adder2()
	    t.Log(f(1), " - ")
	    t.Log(f(28), " - ")
	    t.Log(f(300))
    }
    func Adder2() func(int) int {
        var x int
        return func(delta int) int {
            x += delta
            return x
        }
    }
    函数Adder2() 被赋值到变量f中（类型为func(int) int）
    三次调用函数f的过程中函数Adder2()中的变量delta的值分别为：1，20，300
    从输出可以看出，在多次调用中，变量x的值是被保留的：
    0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321
    闭包函数保存并积累其中变量的值，不管外部函数退出与否，闭包都能够继续操作外部函数中的局部变量。
    这些局部变量同样也可以是参数，例程一中的Adder(a int)

    下面例程清楚展示了Go语言中闭包的使用：
    1）在闭包中使用到的变量可以是在闭包函数内声明也可以是在外部函数中声明的：
    var g int
    go func(i int) {
        s := 0
        for j := 0; j < i; j++ {
            s += j
        }
        g = s
    }(1000)
    这样闭包函数就能够应用到整个集合的元素上，并修改他们的值，然后这些变量就可以计算或者表示全局或平均值

    练习1：使用闭包写斐波那契数列：
    func fib() func() int {
	    a, b := 1, 1
	    return func() int {
		    a, b = b, a+b
		    return b
	    }
    }
    func TestFibonacciClosure(t *testing.T) {
        f := fib()
        for i := 0; i <= 9; i++ {
            t.Log(i, f())
        }
    }
    使用闭包函数作为fib()的返回值

    练习2：理解例程工作原理
    func MakeAddSuffix(suffix string) func(string) string {
        return func(name string) string {
            if !strings.HasSuffix(name, suffix) {
                return name + suffix
            }
            return name
        }
    }
    func TestMakeAddSuffix(t *testing.T) {
        addBmp := MakeAddSuffix(".bmp")
        addJpeg := MakeAddSuffix(".jpeg")
        t.Log(addBmp("file"))
        t.Log(addJpeg("file"))
    }
    首先一个返回值为另一个函数的函数被称为工厂函数，这在需要创建一系列相似函数的时候很有用：
    书写一个工厂函数，其他函数可以调用
    可以返回其他函数的函数和接受其他函数作为参数的函数被称为高阶函数，是函数式语言的特点，函数也是一种值，所以
    Go也具有一些函数式语言的特性，闭包的用法在goroutine和管道中很常用