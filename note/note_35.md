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
    