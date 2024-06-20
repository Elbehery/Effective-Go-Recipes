package main

import (
    "fmt"
)

// cumSum returns the cumulative sum of values
// cumSum([]int{1, 2, 3}) → []int{1, 3, 6}
func cumSum(values []int) []int {
    var cs []int
    s := 0
    for _, val := range values {
        s += val
        cs = append(cs, s)
    }
    return cs
}


func main() {
    fmt.Println(cumSum([]int{1, 2, 3}))
}
