package bloom

import (
    "fmt"
    "testing"
)

func FuzzNumBytes(f *testing.F) {
    f.Add(0) 

    fn := func(t *testing.T, n int) { 
        // Ignore negative numbers
        if n < 0 {
            return
        }
        nBytes := NumBytes(n)
        ok := (nBytes*8 >= n) && ((nBytes-1)*8 <= n) 
        if !ok {
            t.Fatal(nBytes)
        }
    }
    f.Fuzz(fn) 
}


func TestNumBytes(t *testing.T) {
    testCases := []struct {
        nBits  int
        nBytes int
    }{
        {0, 0},
        {1, 1},
        {7, 1},
        {8, 1},
        {9, 2},
    }

    for _, tc := range testCases {
        name := fmt.Sprintf("%+v", tc)
        t.Run(name, func(t *testing.T) {
            if n := NumBytes(tc.nBits); n != tc.nBytes {
                t.Fatal(n)
            }
        })
    }
}
