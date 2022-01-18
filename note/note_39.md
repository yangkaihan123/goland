通过内存缓存来提升程序性能
-
    在程序进行大量计算时，提升性能最直接有效的一种方式就是避免重复计算。
    通过内存中的缓存和重复利用相同计算的结果，称之为内存缓存
    
    拿生成斐波那契数列的程序来说：
    要计算数列中第n个数字，需要先得到之前两数的值，但是很明显绝大多数情况下前两数的值都是计算过的。
    即每个更加后面的数都是基于前面计算的结果重复计算的。
    
    而我们要做的就是将第n个数的值存储在数组中索引为n的位置，然后在数组中查找是否已经计算过，如果没有找到，则再进行计算
    ex：
        const LIM = 41

        var fibs [LIM]uint64

        func TestFibonacciMem(t *testing.T) {
            var result uint64 = 0
            start := time.Now()
            for i := 0; i < LIM; i++ {
                result = fibonacci02(i)
                t.Logf("fibonacci(%d) is %d\n", i, result)
            }
            end := time.Now()
            delta := end.Sub(start)
            t.Logf("fibonacci took this amount of time is %s\n", delta)
        }

        func fibonacci02(n int) (res uint64) {
            if fibs[n] != 0 {
                res = fibs[n]
                return
            }
            if n <= 1 {
                res = 1
            } else {
                res = fibonacci02(n-1) + fibonacci02(n-2)
            }
            fibs[n] = res
            return
        }
    
    内存缓存技术在使用计算成本相对昂贵的函数时候非常有用，比如大量进行相同参数的计算，这种技术还可以用于纯函数中，即相同输入必定获得相同输出的函数
    
