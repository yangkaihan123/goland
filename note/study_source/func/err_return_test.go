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

	t.Log("the second test number is : 5")
	r2, err2 := MySqrt2(5)
	if err2 != nil {
		t.Log("the error! return :", r2, err2)
	} else {
		t.Log("it's ok! return: ", r2, err2)
	}
	t.Log(MySqrt2(5))
}

func MySqrt(f float64) (ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("don't be able to do a sqrt")
	} else {
		ret = math.Sqrt(f)
	}
	return
}

func MySqrt2(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("can't")
	}
	return math.Sqrt(f), nil
}
