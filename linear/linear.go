package main

import "fmt"

func search(haystack *[6]int, needle int) bool {
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle {
            return true
        }
    }
    return false
}

func main() {
    test := [...]int{13, 3, 2, 140, 4, 28}
    fmt.Println("The answer is:", search(&test, 140))
}
