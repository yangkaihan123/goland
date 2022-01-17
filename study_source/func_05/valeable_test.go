package func_05_test

var g int

go fun(i int) {
	s := 0
	for j := 0; j < i; j++ {s += j}
	g = s
}(1000)
