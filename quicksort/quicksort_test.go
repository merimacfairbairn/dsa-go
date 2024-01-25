package main

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	unsorted := []int{8, 2, 10, 3, 5, 6}
	want := []int{2, 3, 5, 6, 8, 10}
	quicksort(unsorted)
	if !reflect.DeepEqual(want, unsorted) {
		t.Errorf("Expected %v, got %v", want, unsorted)
	}
}
