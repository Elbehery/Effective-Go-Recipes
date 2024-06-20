package cart

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func Test_cartTotal(t *testing.T) {
    cart := []LineItem{
        {"lemon", 4, 0.5},
        {"orange", 5, 0.4},
        {"banana", 6, 0.1},
    }
    discounts := map[string]float64{
        "orange": 0.9, // 10% discount on oranges
    }

    expected := 4*0.5 + 5*0.4*0.9 + 6*0.1
    total := cartTotal(cart, discounts)
    require.InDelta(t, expected, total, 0.001)
}

