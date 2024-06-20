package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

type workerInfo struct {
    Info
    err error
}


type infoReq struct {
    url string            // URL to query
    ch  chan<- workerInfo // return channel
}


func infoWorker(ch <-chan infoReq) {
    for req := range ch {
        log.Printf("info request for: %s", req.url)
        info, err := siteInfo(req.url)
        log.Printf("%s: %#v (err=%v)", req.url, info, err)
        winfo := workerInfo{
            Info: info,
            err:  err,
        }
        req.ch <- winfo
    }
}



// Pool is a fixed pool of workers.
type Pool struct {
    queue chan infoReq
}

// NewPool creates a new Pool with n workers.
func NewPool(n int) (*Pool, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0 (got %d)", n)
    }

    queue := make(chan infoReq)
    for i := 0; i < n; i++ {
        go infoWorker(queue)
    }

    p := Pool{
        queue: queue,
    }

    return &p, nil
}

// Close signals the worker goroutines to terminate.
func (p *Pool) Close() error {
    if p.queue != nil {
        close(p.queue)
        p.queue = nil
    }
    return nil
}



// SiteInfo returns a channel with info.
func (p *Pool) SiteInfo(url string) (Info, error) {
    // Return channel, buffered to avoid goroutine leak
    ch := make(chan workerInfo, 1) 
    p.queue <- infoReq{url, ch}
    info := <-ch
    return info.Info, info.err
}



// Info is information on site.
type Info struct {
    statusCode int
    delay      time.Duration
}

func siteInfo(url string) (Info, error) {
    var info Info

    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        return info, err
    }
    defer resp.Body.Close()

    // Consume data from server
    _, err = io.Copy(io.Discard, resp.Body)
    if err != nil {
        return info, err
    }

    info.delay = time.Since(start)
    info.statusCode = resp.StatusCode
    return info, nil
}


func main() {
    urls := []string{
        "https://www.apple.com/",
        "https://www.microsoft.com/",
        "https://www.ibm.com/",
        "https://www.dell.com/",
    }

    pool, err := NewPool(4)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer pool.Close()

    for _, url := range urls {
        info, err := pool.SiteInfo(url)
        fmt.Printf("%#v (err=%v)\n", info, err)
    }
}
