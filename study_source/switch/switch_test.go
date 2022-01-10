package switch_test

import "testing"

func TestSwitch(t *testing.T) {
	num1 := 100
	switch num1 {
	case 98, 99:
		t.Log("it's equal to 98")
	case 100:
		t.Log("it's equal to 100")
	default:
		t.Log("it's not equal to 98 or 99")
	}
}

func TestSwitch2(t *testing.T) {
	num1 := 7
	switch {
	case num1 < 0:
		t.Log("number is negative")
	case num1 > 0 && num1 < 10:
		t.Log("number is between 0 and 10")
	default:
		t.Log("number is 10 or greater")
	}
}

func TestSeason(t *testing.T) {
	var a int = 2
	switch {
	case a >= 1 && a <= 3:
		t.Log("the season is spring")
	case a >= 4 && a <= 6:
		t.Log("the season is summer")
	case a >= 7 && a <= 9:
		t.Log("the season is autumn")
	case a >= 10 && a >= 12:
		t.Log("the season is winter")
	default:
		t.Log("no season")
	}
}
