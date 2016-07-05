package linkedList

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	nums := []int{1, 2, 2, 3, 4, 4, 5}
	output := []int{1, 2, 3, 4, 5}

	ll := NewNode(nil)
	for _, n := range nums {
		ll.Insert(Int(n))
	}

	ll.RemoveDuplicates()

	equal := reflect.DeepEqual(ll.Traverse(), output)
	if !equal {
		t.Error("Removing duplicate failed")
	}

}
