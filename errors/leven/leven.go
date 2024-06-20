package main

import (
    "fmt"
)

// emulate external package
var dist struct {
    Edit func(s1, s2 string) int
}

func init() {
    dist.Edit = func(s1, s2 string) int {
        c1 := s1[10]
        return int(c1)
    }
}


// EditDistance Returns the edit (Levenshtein) distance  between s1 and s2.
// It wraps dist.Edit against panics.
func EditDistance(s1, s2 string) (distance int, err error) {

    defer func() {
        if e := recover(); e != nil {
            err = fmt.Errorf("%v", e) // Convert e (any) to error
        }
    }()

    return dist.Edit(s1, s2), nil
}

func main() {
    fmt.Println(EditDistance("kitten", "sitting"))
}
