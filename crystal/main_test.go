package main

import "testing"

func TestSearch(t *testing.T) {
    test := [...]bool{false, false, false, false, true, true}
    want := 4
    result := crystal(&test)
    if result != want {
        t.Errorf("Expected %v, got %v for the array %v", want, result, test)
    }
}
