package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Connect to database ...
    fmt.Fprintf(w, "OK\n")
}

func main() {
    http.HandleFunc("/health", healthHandler)

    addr := os.Getenv("HTTPD_ADDR")
    if addr == "" {
        addr = ":8080"
    }
    log.Printf("sever listens on %s", addr)

    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
