package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)
    close(ch)

    val, ok := <-ch
    fmt.Println("val =", val, "ok =", ok)

    var i any = "hi"
    // n := i.(int) will panic
    if n, ok := i.(int); ok {
        fmt.Println("int", n)
    } else {
        fmt.Println("not an int")
    }
}
