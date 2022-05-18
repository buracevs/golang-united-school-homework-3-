package homework

import (
	"reflect"
	"testing"
)

func TestSortMapValues(t *testing.T) {
	initial := map[int]string{2: "a", 0: "b", 1: "c"}
	expected := []string{"b", "c", "a"}

	actual := sortMapValues(initial)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Slices are not equal, expected:%v but was %v", expected, actual)
	}
}
