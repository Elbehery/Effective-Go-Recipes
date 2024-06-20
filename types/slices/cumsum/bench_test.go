package main

import (
    "testing"
)

var (
    values   []int
    size     = 9323
    expected int
)

func init() {
    for i := 0; i < size; i++ {
        values = append(values, i+1)
    }
    // sum of odd numbered arithmetic series
    expected = size * ((values[0] + values[size-1]) / 2)
}


func TestCumsum(t *testing.T) {
    out := cumSum(values)
    last := out[len(out)-1]
    if expected != last {
        t.Fatalf("%d != %d", expected, last)
    }
}

func BenchmarkCumsum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        out := cumSum(values) // values is populated in init
        // Make sure compiler doesn't optimize the above call away
        if len(out) != len(values) {
            b.Fatalf("bad out length: %d", len(out))
        }
    }
}

