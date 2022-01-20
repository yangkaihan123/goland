package arr

import "testing"

func TestArraysum(t *testing.T) {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // note the explicit address-of operator
	//to pass a pointer to the array
	t.Logf("the sum of the array is %f,", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		//derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
