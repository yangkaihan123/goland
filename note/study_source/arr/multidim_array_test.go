package arr

import (
	"testing"
)

const (
	WIDTH  = 10
	HEIGHT = 5
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TestMultidim_array(t *testing.T) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	t.Log(screen)
}
