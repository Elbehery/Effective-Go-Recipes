package main

import (
    "flag"
    "fmt"
)

var (
    version = "<unknown>"
)


func main() {
    var showVersion bool
    flag.BoolVar(&showVersion, "version", false, "show version and exit")
    flag.Parse()

    if showVersion {
        fmt.Printf("agent version %s\n", version)
    }
}
