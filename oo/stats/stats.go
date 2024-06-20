package stats

import (
    "fmt"
)


// Number is set of possible numbers.
type Number interface {
    ~float64 | ~int
}



// Max returns the maximal value in values.
func Max[T Number](values []T) (T, error) {
    if len(values) == 0 {
        var zero T 
        return zero, fmt.Errorf("Max of empty slice")
    }

    max := values[0]
    for _, v := range values[1:] {
        if v > max {
            max = v
        }
    }
    return max, nil
}

