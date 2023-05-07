package patterns

import (
	"reflect"
	"testing"
)

func TestIterator(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	i := NewIterator[int](items)

	var got []int

	for i.HasNext() {
		got = append(got, i.GetNext())
	}

	if !reflect.DeepEqual(got, items) {
		t.Errorf("iterator = %v, want %v", got, items)
	}
}
