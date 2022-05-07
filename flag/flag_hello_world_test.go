package flag

import (
	"flag"
	"fmt"
	"testing"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everone", "the greeting object")
}
func TestFlag(t *testing.T) {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
