package func03_test

import (
	"errors"
	"math"
	"testing"
)

func TestErrReturn(t *testing.T) {
	t.Log("the first test number is : -1")
	ret1, err1 := MySqrt(5)
	if err1 != nil {
		t.Log("the error! return values are: ", ret1, err1)
	} else {
		t.Log("it's ok! return values are: ", ret1, err1)
	}
}

func MySqrt(f float64) (ret float64, err error) {
	if f < 0 {
		ret = math.NaN()
		err = errors.New("don't be able to do a sqrt")
	} else {
		ret = math.Sqrt(f)
	}
	return
}
