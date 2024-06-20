package geo

import (
    "math"
)

func Euclidean(x1, y1, x2, y2 float64) float64 {
    dx := x1 - x2
    dy := y1 - y2

    // Hypot won't under/overflow
    return math.Hypot(dx, dy)
}

