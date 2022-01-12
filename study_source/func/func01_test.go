package func02_test

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	t.Log("In main before calling greeting")
	greeting()
	t.Log("in main after calling greeting")
}

func greeting() {
	fmt.Println("heelless")
}
