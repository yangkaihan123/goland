传递变长参数
-
    在函数里，如果函数的最后一个参数是采用 ...type的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为0
    这样的函数被称为变参函数
    func myFunc(a,b,arg ...int) {}
    这个函数，接受一个类似某个类型的slice参数，切片参数可以用for循环迭代
    
    ex:
        func Greeting(prefix string, who ...string)
        Greeting("hello:","Joe","Anna","Eileen")
    例程中，Greeting函数中，变量who的值为[]string{"Joe","Anna","Eileen"}
    如果参数被存储在一个切片里，可以通过slice... 的形式来传参，调用变参函数
    ex2:
        func TestCanChange(t *testing.T) {
            x := min(1, 3, 2, 0)
	        t.Logf("the minfunc is %d\n", x)
	        b := []int{7, 9, 3, 5, 1}
	        a := min(b...)
	        t.Logf("the minfunc is %d\n", a)
        }

        func min(s ...int) int {
            if len(s) == 0 {
                return 0
            }
            min := s[0]
            for _, v := range s {
                if v < min {
                    min = v
                }
            }
            return min
        }
    例程中：a在TestCanChange中作为min的返回值，传入的参数为b（一个切片），就需要用 b... 来传参

    变长参数可以作为对应类型的切片slice进行二次传递
    当变长参数的类型不是都相同的时候有两种方式：
    （1）使用结构体：
        type Options struct {
            par1 type1,
            par2 type2,
            ...
        }
        定义一个结构类型为Options，用来存储所有可能的参数
        函数F1可以正常使用参数a和b，以及一个没有任何初始化的Options结构：
        F1(a, b, Options {})
        如果需要对选项进行初始化，可以使用：
        F1(a, b, Options {par1:var1, par2:var2})
    （2）如果一个变长参数的类型没有被指定，可以使用默认的空接口（interface）可以接受任何类型的参数
        这个方式不仅可以用于未知长度的参数，还可以用于任何不确定类型的参数。一般会使用一个for-range循环，或者switch
        对每个参数的类型进行判断。
        ex:
            for typecheck(..,..,values ... interface{}) {
                for _, value := range values {
                    switch v := values.(type) {
                        case int: ...
                        case float: ...
                        case string: ...
                        case bool: ...
                        default: ...
                    }
                }
            }