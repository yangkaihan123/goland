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
    