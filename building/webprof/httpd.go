package main

import (
    "fmt"
    "log"
    "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Run health checks
    fmt.Fprintf(w, "OK\n")
}

func main() {
    http.HandleFunc("/health", healthHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
