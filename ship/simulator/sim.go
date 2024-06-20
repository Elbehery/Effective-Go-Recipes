package sim

import (
    "github.com/353solutions/geo"
)

func Distance(x1, y1, x2, y2 float64) float64 {
    return geo.Euclidean(x1, y1, x2, y2)
}
