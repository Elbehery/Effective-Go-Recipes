package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "net/http/httptest"
    "sync"
    "sync/atomic"
    "time"
)

var (
    callCount int64
)


func handler(w http.ResponseWriter, r *http.Request) {
    atomic.AddInt64(&callCount, 1)

    // Simulate work
    time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
    w.Write([]byte("OK"))
}


func main() {
    const nRuns = 1000
    const nGoroutines = 10

    var wg sync.WaitGroup
    wg.Add(nGoroutines)
    for g := 0; g < nGoroutines; g++ {
        go func() {
            defer wg.Done()
            for i := 0; i < nRuns; i++ {
                // Dummy ResponseWriter & http.Request for the handler
                w := httptest.NewRecorder()
                r := httptest.NewRequest("GET", "localhost:8080", nil)
                handler(w, r)
            }
        }()
    }

    wg.Wait()
    fmt.Println(callCount)
}
