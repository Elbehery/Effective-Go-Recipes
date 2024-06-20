package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

type Metric struct {
    Name  string    `json:"name"`
    Host  string    `json:"host"`
    Time  time.Time `json:"time"`
    Value float64   `json:"value"`
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    dec := json.NewDecoder(r.Body)
    count := 0
    for {
        var m Metric
        err := dec.Decode(&m)
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("error: %s", err)
        }
        log.Printf("metric: %+v", m)
        count++
    }

    log.Printf("%d metrics\n", count)
    fmt.Fprintf(w, "%d metrics\n", count)
}

func main() {
    http.HandleFunc("/metrics", metricsHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
