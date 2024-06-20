package main

import (
    "fmt"
)

func main() {
    currCap := 0
    var s []int
    for i := 0; i < 2000; i++ {
        s = append(s, i)
        if c := cap(s); c != currCap {
            ratio := float64(c) / float64(currCap)
            fmt.Printf("%4d → %4d (%.2f)\n", currCap, c, ratio)
            currCap = c
        }
    }
}

