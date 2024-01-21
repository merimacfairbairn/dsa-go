package main

import (
	"fmt"
	"math"
)

func crystal(breaks *[6]bool) int {
    jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

    i := jmpAmount
    for ; i < len(breaks); i += jmpAmount {
        if breaks[i] {
            break
        }
    }

    i -= jmpAmount

    for j := 0; j <= jmpAmount && j < len(breaks); j, i = j+1, i+1 {
        if breaks[i] {
            return i
        }
    }
    return -1
}

func main() {
    test := [...]bool{false, false, false, false, true, true}
    fmt.Println("The answer is:", crystal(&test))
}
