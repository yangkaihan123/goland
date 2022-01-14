递归函数
-
    当一个函数在其函数体内调用自身，称之为递归。
    最经典的例子就是计算斐波拉契数列：
    前两个数为1，从第三个数开始每个数均为前两个数之和
    1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
    Fibonacci实现例程：
    func fibonacci(n int) (res int) {
	    if n <= 1 {
		    res = 1
	    } else {
		    res = fibonacci(n-1) + fibonacci(n-2)
	    }
	    return
    }
    func TestFibonacci(t *testing.T) {
        result := 0
        for i := 0; i <= 10; i++ {
            result = fibonacci(i)
            t.Logf("fibonacci(%d) is: %d\n", i, result)
        }
    }
    实现了10个Fibonacci数

    还有很多问题都可以使用递归来解决
    但是使用递归函数时会常遇到一个重要问题就是栈溢出，一般出现在大量递归调用时导致程序栈内存分配耗尽，这个问题
    可以通过一个名为懒惰求值的方法解决
    
    go语言中也可以使用相互调用的递归函数，多个函数之间相互调用形成闭环，这些函数声明顺序时任意的。
    例程：
        func TestDigui(t *testing.T) {
	        t.Logf("%d is even: is %t\n", 16, even(16))
        }
        func even(nr int) bool {
            if nr == 0 {
                return true
            }
            return odd(RevSign(nr) - 1)
        }
        func RevSign(nr int) int {
            if nr < 0 {
                return -nr
            }
            return nr
        }
        func odd(nr int) bool {
            if nr == 0 {
                return false
            }
            return even(RevSign(nr) - 1)
        }
    例程展示了函数odd和even的关系

    fibonacci数列的另一种实现方式：
    func fibonacci2(n int) (val, pos int) {
        if n <= 1 {
            val = 1
        } else {
            v1, _ := fibonacci2(n - 1)
            v2, _ := fibonacci2(n - 2)
            val = v1 + v2
        }
        pos = n
        return
    }
    func TestFibonacci02(t *testing.T) {
        pos := 5
        res, pos := fibonacci2(pos)
        t.Logf("the %d-th Fibonacci number is %d\n", pos, res)
        pos = 10
        res, pos = fibonacci2(pos)
        t.Logf("the %d-th Fibonacci number is %d\n", pos, res)
    }
    
    练习2：1-10的递归
    func TestRecursive10To1(t *testing.T) {
	    printrec(1)
    }
    func printrec(i int) {
        if i > 10 {
        return
        }
        printrec(i + 1)
        fmt.Printf("%d\n", i)
    }

    练习3：实现一个输出前30个整数的阶乘的程序
    n! 阶乘的定义：n! = n * (n-1)!, 0! = 1
    func Factorial(n uint64) (fac uint64) {
	    fac = 1
	    if n > 0 {
		    fac = n * Factorial(n-1)
		    return
	    }
	    return
    }
    func TestFactorial(t *testing.T) {
        for i := uint64(0); i < uint64(30); i++ {
        t.Logf("Factorial of %d is %d\n", i, Factorial(i))
        }
    }
    
    /* unnamed return variables: 返回未命名返回值
    func Factorial_02(n uint64) uint64 {
        if n > 0 {
        return n * Factorial(n-1)
        }
        return 1
    }
    */
    特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误
    最好的解决方案是使用big包
