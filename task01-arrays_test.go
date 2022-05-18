package homework

import "testing"

func TestAverage(t *testing.T) {
	slice := [15]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}

	actual := average(slice)
	if actual != 3 {
		t.Error("Expected:3, actual was:", actual)
	}
}
