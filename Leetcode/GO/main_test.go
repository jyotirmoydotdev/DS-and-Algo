package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1  []int
		list2  []int
		merged []int
	}{
		{
			list1:  []int{1, 3, 5},
			list2:  []int{2, 4, 6},
			merged: []int{1, 2, 3, 4, 5, 6},
		}, {
			list1:  []int{1, 2, 4},
			list2:  []int{1, 3, 4},
			merged: []int{1, 1, 2, 3, 4, 4},
		}, {
			list1:  []int{},
			list2:  []int{0},
			merged: []int{0},
		},
	}
	for _, test := range tests {
		list1 := createList(test.list1)
		list2 := createList(test.list2)

		result := mergeTwoLists(list1, list2)
		resultSlice := listToSlice(result)

		if !reflect.DeepEqual(resultSlice, test.merged) {
			t.Errorf("Expected %v, got %v", test.merged, resultSlice)
		}
	}
}
