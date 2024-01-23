package main

func search(haystack *[6]int, needle int) bool {
    lo := 0
    hi := len(*haystack)

    for {
        var m int = lo + (hi - lo) / 2
        v := (*haystack)[m]
        if v == needle {
            return true
        }else if v > needle {
            hi = m
        }else {
            lo = m + 1
        }

        if lo >= hi {
            break
        }
    }

    return false
}

func main() {
}
