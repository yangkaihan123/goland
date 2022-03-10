package main

func test04() (int, int) {
	return 1, 2
}

func sum01(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func main() {
	println(add(test04()))
	println(sum01(test04()))
}
