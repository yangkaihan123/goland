计算函数执行时间
-
    有时候，可以知道一个计算执行消耗的时间是很有必要的，比如在对比和基准测试中。
    最简单的一个办法就是在计算开始之前设置一个起始值，再由计算结束时的结束值，最后取出它们之间的差值，就是这个计算所消耗的时间
    可以使用time包中的Now() 和 Sub 函数：
    start := time.Now()
    longCalculation()
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("longCalculation took this amount of time: %s\n",delta)
    
    使用Fibonacci函数来体现会更明显：
    func TestRuntime(t *testing.T) {
	    result := 0
	    start := time.Now()
	    for i := 0; i <= 25; i++ {
		    result = fibonacci(i)
		    t.Logf("fibonacci(%d) is %d\n", i, result)
	    }
	    end := time.Now()
	    delta := end.Sub(start)
	    t.Logf("fibonacci took this amount of time: %s\n", delta)
    }
    func fibonacci(n int) (res int) {
        if n <= 1 {
            res = 1
        } else {
            res = fibonacci(n-1) + fibonacci(n-2)
        }
        return
    }

    输出结果为：
    === RUN   TestRuntime
    runtime_test.go:13: fibonacci(0) is 1
    runtime_test.go:13: fibonacci(1) is 1
    runtime_test.go:13: fibonacci(2) is 2
    runtime_test.go:13: fibonacci(3) is 3
    runtime_test.go:13: fibonacci(4) is 5
    runtime_test.go:13: fibonacci(5) is 8
    runtime_test.go:13: fibonacci(6) is 13
    runtime_test.go:13: fibonacci(7) is 21
    runtime_test.go:13: fibonacci(8) is 34
    runtime_test.go:13: fibonacci(9) is 55
    runtime_test.go:13: fibonacci(10) is 89
    runtime_test.go:13: fibonacci(11) is 144
    runtime_test.go:13: fibonacci(12) is 233
    runtime_test.go:13: fibonacci(13) is 377
    runtime_test.go:13: fibonacci(14) is 610
    runtime_test.go:13: fibonacci(15) is 987
    runtime_test.go:13: fibonacci(16) is 1597
    runtime_test.go:13: fibonacci(17) is 2584
    runtime_test.go:13: fibonacci(18) is 4181
    runtime_test.go:13: fibonacci(19) is 6765
    runtime_test.go:13: fibonacci(20) is 10946
    runtime_test.go:13: fibonacci(21) is 17711
    runtime_test.go:13: fibonacci(22) is 28657
    runtime_test.go:13: fibonacci(23) is 46368
    runtime_test.go:13: fibonacci(24) is 75025
    runtime_test.go:13: fibonacci(25) is 121393
    runtime_test.go:17: fibonacci took this amount of time: 1.1024ms
    --- PASS: TestRuntime (0.00s)
    PASS