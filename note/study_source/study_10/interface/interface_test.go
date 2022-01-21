package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
} // 使用duck type，把Programmer这个interface 绑定到GoProgrammer struct

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
