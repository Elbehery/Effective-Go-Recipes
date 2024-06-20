package sim_test

import (
    "math"
    "testing"

    "github.com/353solutions/geo"
)

func TestEuclideanBug(t *testing.T) {
    d := geo.Euclidean(0, 0, 95e200, 168e200)
    if math.IsInf(d, 1) {
        t.Fatal(d)
    }
}

