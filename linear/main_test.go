package main

import "testing"

func TestSearch(t *testing.T) {
    haystack := [...]int{12, 3, 3, 22, 10, 874}
    needle := 3
    want := true
    result := search(&haystack, needle)
    if result != want {
        t.Errorf("Expected %v, got %v for the array %v", want, result, haystack)
    }
}
