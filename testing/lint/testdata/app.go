package main

import (
    "fmt"
    "syscall" // want `importing forbidden package "syscall"`
)

func main() {
    fmt.Println(syscall.AF_ALG)
}
