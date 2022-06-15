package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object")
}

func main() {
	flag.Parse()
	hello(name)
}
