package main

import (
    "fmt"
    "log"
    "net/http"
)

func crashHandler(w http.ResponseWriter, r *http.Request) {
    go func() {
        panic("down we go!")
    }()
    fmt.Fprintln(w, "OK")
}

func main() {
    http.HandleFunc("/carsh", crashHandler)

    addr := ":8080"
    log.Printf("server starting on %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
