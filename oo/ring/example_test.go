package stats

import (
    "fmt"
)

func ExampleRing() {
    size := 4
    r, err := NewRing[int](size)
    if err != nil {
        fmt.Printf("error: %s", err)
        return
    }

    for i := 1; i <= 10; i++ {
        r.Push(i)
    }
    fmt.Println(r.Mean())

    // Output:
    // 8.5
}

