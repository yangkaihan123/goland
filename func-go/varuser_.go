package main

func test_03() (int, int) {
	return 1, 2
}

func main() {
	// s := make([]int,2)
	// s = test_03() // Error: multiple-value test_03() in single-value context

	x, _ := test_03()
	println(x)
}

/*
	_ 占位符 忽略了第二个返回值
*/
