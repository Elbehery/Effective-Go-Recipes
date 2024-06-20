package cart


// LineItem is an item in a shopping cart.
type LineItem struct {
    Name   string
    Amount float64 // units or weight in pounds
    Price  float64
}



// cartTotal calculates total price of cart after applying discounts.
// discounts is a map from Name -> discount value.
func cartTotal(cart []LineItem, discounts map[string]float64) float64 {
    total := 0.0
    for _, li := range cart {
        discount := discounts[li.Name]
        total += li.Amount * li.Price * discount
    }
    return total
}

