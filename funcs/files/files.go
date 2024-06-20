package main

import (
    "fmt"
    _ "unsafe"

    "git.corp.com/client"
)

// HACK: Access internal client.setMaxBodySize.
// Remove this once setMaxBodySize becomes exported, see issue #732.

//go:linkname setClientMaxBodySize git.corp.com/client.setMaxBodySize
func setClientMaxBodySize(int64)


func main() {
    setClientMaxBodySize(5 * 1_000_000) // 5 MB
    c := client.New()
    // Start using client with big images...
    fmt.Println(c) // Make compiler happy
}
