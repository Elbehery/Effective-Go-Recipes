package bloom


// NumBytes returns the number of bytes that can hold n bits.
func NumBytes(nBits int) int {
    return (nBits + 7) / 8
}

