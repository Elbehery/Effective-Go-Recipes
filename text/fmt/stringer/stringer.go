package main

import (
    "fmt"
    "time"
)

func formatValue(val any) string {
    if s, ok := val.(fmt.Stringer); ok { // val implements fmt.Stringer
        return s.String()
    }

    // Return default formatting ...
    switch val.(type) {
    // TODO ...
    }
    return fmt.Sprintf("%v", val)
}

func main() {
    t := time.Now()
    fmt.Println(formatValue(t))
}
