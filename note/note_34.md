将函数作为参数
-
    函数可以作为其它函数的参数进行传递，然后在其他函数内调用执行，被称为回调
    ex:
        func TestCallBack(t *testing.T) {
	        callback(1, Add)
        }
        func Add(a, b int) {
            fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
        }
        func callback(y int, f func(int, int)) {
            f(y, 2) //this becomes Add(1,2)
        }
    例程输出：The sum of 1 and 2 is: 3

    将函数作为参数最佳实践是函数 strings.IndexFunc():
    这个函数的签名是：
    func IndexFunc(s string, f func(c int) bool) int
    返回值是在函数f(c)返回true，-1或者从未返回时的索引值
    例如：strings.IndexFunc(line, unicode.IsSpace)
         就会返回line 中第一个空白字符的索引值
    也可以写自己的函数：
        func IsAscii(c int) bool {
            if c < 255 {
                return false
            }
            return true
        }
    
    后面会根据一个客户端/服务端程序作为例程对这个用法进行研究
    type binOp func(a, b int) int
    func run(op binOp, req *Request) { ... }

    练习程序：
    将指定文本内的所有非 ASCII 字符替换成 ? 或空格
    func TestStringsMap(t *testing.T) {
	    asciiOnly := func(c rune) rune {
	    	if c > 127 {
			    return ' '
		    }
		    return c
	    }
	    t.Log(strings.Map(asciiOnly, "Jérôme Österreich"))
    }
    使用包strings中的Map函数实现
    