package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "math/rand"
    "net/http"
    "os"
    "time"
)

var (
    metrics   = []string{"CPU", "Memory", "Disk"}
    hosts     = []string{"srv1", "srv2", "srv3", "srv4", "srv5"}
    serverURL = "http://localhost:8080/metrics"
)

func collectMetrics() <-chan Metric {
    ch := make(chan Metric)
    go func() {
        for i := 0; i < 20; i++ {
            m := Metric{
                Name:  metrics[i%len(metrics)],
                Host:  hosts[i%len(hosts)],
                Time:  time.Now(),
                Value: rand.Float64() * 100,
            }
            ch <- m
            time.Sleep(117 * time.Millisecond)
        }
        close(ch)
    }()

    return ch
}


// Metric is a metric in the system.
type Metric struct {
    Name  string    `json:"name"`
    Host  string    `json:"host"`
    Time  time.Time `json:"time"`
    Value float64   `json:"value"`
}


func producer(w io.WriteCloser) {
    defer w.Close()

    enc := json.NewEncoder(w)
    for m := range collectMetrics() {
        if err := enc.Encode(m); err != nil {
            log.Printf("error: can't encode %#v - %s", m, err)
            return
        }
    }
}


func updateMetrics() error {
    r, w, err := os.Pipe()
    if err != nil {
        return err
    }

    go producer(w)

    req, err := http.NewRequest("POST", serverURL, r)
    if err != nil {
        return err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    if resp.StatusCode != http.StatusOK {
        const format = "bad reply status: %d %s"
        return fmt.Errorf(format, resp.StatusCode, resp.Status)
    }

    return nil
}


func main() {
    if err := updateMetrics(); err != nil {
        log.Fatalf("error: %s", err)
    }
}
