package main

import (
    "fmt"
    "io"
    "strings"
    "unicode"
)

func main() {
    data := `
We have 1234
then 2342
then 110
then 37
`
    var df DigitsFreq 
    io.Copy(&df, strings.NewReader(data))
    for r, c := range df.Freqs {
        fmt.Printf("%c → %d\n", r, c)
    }
}


// Write implements io.Writer.
func (d *DigitsFreq) Write(data []byte) (int, error) {
    if d.Freqs == nil { 
        d.Freqs = make(map[rune]int)
    }

    for _, b := range data {
        if r := rune(b); unicode.IsDigit(r) {
            if !d.inNum { // it's a leading digit
                d.Freqs[r]++
                d.inNum = true
            }
            continue
        }

        // Here it's not a digit
        if d.inNum {
            d.inNum = false
        }
    }

    return len(data), nil
}



// DigitsFreq calculates leading digit frequency.
type DigitsFreq struct {
    Freqs map[rune]int // Leading digit frequency

    inNum bool // local state
}

