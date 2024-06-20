package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// addLogging is a middleware that adds logging to a handler.
func addLogging(name string, handler http.Handler) http.Handler {
    wrapper := func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        handler.ServeHTTP(w, r)

        duration := time.Since(start)
        log.Printf("%s took %s", name, duration)
    }

    return http.HandlerFunc(wrapper)
}


func queryHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TBD\n")
}

func main() {
    hdlr := addLogging("query", http.HandlerFunc(queryHandler))
    http.Handle("/query", hdlr)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
