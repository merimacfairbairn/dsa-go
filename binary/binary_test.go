package main

import "testing"

func TestSearch(t *testing.T) {
    haystack := [...]int{12, 39, 123, 98, 659, 3}
    needle := 98
    want := true
    result := search(&haystack, needle)

    if result != want {
        t.Errorf("Expected %v, got %v for arr %v", want, result, haystack)
    }
}
