package stats

import (
    "fmt"
)

func ExampleMax() {
    iVals := []int{15, 42, 16, 8, 23, 4}
    fmt.Println(Max(iVals))

    fVals := []float64{3.14, 2.718, 6.283, 1.618}
    fmt.Println(Max(fVals))

    _, err := Max[int](nil) 
    fmt.Println(err)

    // Output:
    // 42 <nil>
    // 6.283 <nil>
    // Max of empty slice
}

