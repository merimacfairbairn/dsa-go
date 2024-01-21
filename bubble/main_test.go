package main

import "testing"

func TestBubble(t *testing.T) {
    test := [...]int{31, 38, 30, 3, 62, 1}
    bubble(&test)
    want := [...]int{1, 3, 30, 31, 38, 62}
    if test != want {
        t.Errorf("Expected %v, got %v", want, test)
    }
}
