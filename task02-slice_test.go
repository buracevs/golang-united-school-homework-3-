package homework

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	initial := []int64{0, 1, 2}
	actual := reverse(initial)
	expected := []int64{2, 1, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Slices are not equal, expected:%v but was %v", expected, actual)
	}
}
