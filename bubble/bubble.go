package main

import (
	"fmt"
)

func bubble(arr *[6]int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    test := [...]int{3, 8, 0, 5, 6, 1}
    bubble(&test)
    fmt.Println("Sorted array looks like:", test)
}
