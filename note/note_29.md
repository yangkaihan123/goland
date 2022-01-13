函数参数与返回值
-

###函数的参数
    函数能够接收参数供自己使用， 也可以返回零个或多个值（通常把返回多个值称为返回一组值）
    多返回值是Go的一大特性，可以用来判断一个函数是否正常执行

    使用return关键字返回一组值，任何一个有返回值的函数都必须以return或者panic来结尾
    在函数块里return之后的语句都不会执行，如果一个函数需要返回值，那么这个函数里的每一个代码分支（code-path）都要有return

    ex: 这个函数不能被编译 这里得研究研究 还不明白
        func (st *Stack) Pop() int {
            v := 0
            for ix := len(st) - 1; ix >= 0; ix-- {
                if v = st[ix]; v != 0 {
                    st[ix] = 0
                    return v
                }
            }
        }
    函数定义时形参一般是有名字的，也可以定义没有形参名的函数，只有相应的形参类型：func f(int, int, float64)
    没有参数的函数通常被称为niladic函数，类似的 main.main()

###按值传递（call by value）和按引用传递（call by reference）
    Go默认使用按值传递来传参，也就是传递参数的副本（副本）
    函数接收参数副本以后，在使用变量的过程中可能对副本的值进行更改，但是不会影响原值
    比如: Function(arg1)
    如果想让函数直接修改参数值，而不是对参数的副本进行操作，需要将参数地址（其实就是指针了，变量名前加&）传递给函数，这样就是按引用传递
    比如：Function(&arg1)，此时传递给函数的是一个指针
    如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但是指针上的值所指向的地址上的值不会被复制（就是源值）
    可以通过修改指针的值来修改指向这个值所指向的地址上的值
    （指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址，所以按引用传递的本质也是按值传递）
    因为传递指针的消耗比传递副本会少，所以在函数调用时，像切片（slice）字典（map）接口（interface）通道（channel）
    这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
    有些函数只是完成一个任务，并没有返回值，仅仅为了利用这个函数的作用来打印文本，发送邮件，记录panic等一些函数，但是绝大部分函数都是有明确返回值的

    在example：func02_test.go中
    MultiPly3Nums函数带有三个形参，分别是a,b,c 还有一个int类型的返回值
    func TestFunc02(t *testing.T) {
	t.Logf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
    }
    func MultiPly3Nums(a int, b int, c int) int {
        return a * b *c
    }
    
    当函数需要多返回值，四个到五个，可以传递一个切片给函数，前提是返回值具有相同类型，当返回值有不同类型时候，可以传递一个结构体
    Q：下面两个函数调用的区别：
    （A）func DoSomething(a *A) {
            b = a
        }

    （B）func DoSomething(a A) {
            b = &a
        }
    一个是值传递一个是引用传递

###命名的返回值（named return variables)
   
    var (
	num   int = 10
	numx2 int
	numx3 int
    )
    func TestFunctionReturn(t *testing.T) {
        numx2, numx3 = getX2AndX3(num)
        PrintValues()
        numx2, numx3 = getx2andx32(num)
        PrintValues()
    }
    func getx2andx32(input int) (x2 int, x3 int) {
        x2 = 2 * input
        x3 = 3 * input
        return
    }
    func PrintValues() {
        fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
    }
    func getX2AndX3(input int) (int, int) {
        return 2 * input, 3 * input
    }
    这个例程中，getx2andx32和getX2AndX3都是带有一个int参数，返回两个int值
    其中一个函数的返回值在函数调用时就已经被赋予了一个初始零值
    getX2AndX3 与 getX2AndX3_2 两个函数演示了如何使用非命名返回值与命名返回值的特性。
    当需要返回多个非命名返回值时，需要使用 () 把它们括起来，比如 (int, int)
    命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，只需要一条简单的不带参数的 return 语句。
    需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来，PrintValues()

    警告：
        。return 或 return var都是可以的
        。但是return var = expression（表达式）会引发一个编译错误

    即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。
    任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）
    尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂