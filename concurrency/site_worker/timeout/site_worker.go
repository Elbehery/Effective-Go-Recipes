package main

import (
    "context"
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
    ctx context.Context
    url string            // URL to query
    ch  chan<- workerInfo // return channel
}


func infoWorker(ch <-chan infoReq) {
    for req := range ch {
        log.Printf("info request for: %s", req.url)
        outCh := make(chan workerInfo, 1)
        go func() { 

            info, err := siteInfo(req.ctx, req.url)
            log.Printf("%s: %#v (err=%v)", req.url, info, err)
            winfo := workerInfo{
                Info: info,
                err:  err,
            }
            outCh <- winfo
        }()

        select {
        case info := <-outCh:
            log.Printf("%s: %#v", req.url, info)
            req.ch <- info
        case <-req.ctx.Done():
            req.ch <- workerInfo{err: req.ctx.Err()}
        }
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



// SiteInfo return a channel with info.
func (p *Pool) SiteInfo(ctx context.Context, url string) (Info, error) {
    // Return channel, buffered to avoid goroutine leak
    ch := make(chan workerInfo, 1)

    // Send timeout
    select { 
    case p.queue <- infoReq{ctx, url, ch}:
        // Nothing to do here
    case <-ctx.Done():
        return Info{}, ctx.Err()
    }

    // Receive timeout
    select { 
    case info := <-ch:
        return info.Info, nil
    case <-ctx.Done():
        return Info{}, ctx.Err()
    }
}



// Info is information on site.
type Info struct {
    statusCode int
    delay      time.Duration
}

func siteInfo(ctx context.Context, url string) (Info, error) {
    var info Info

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return Info{}, nil
    }

    start := time.Now()
    resp, err := http.DefaultClient.Do(req)
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

    timeout := time.Second
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    for _, url := range urls {
        info, err := pool.SiteInfo(ctx, url)
        fmt.Printf("%#v (err=%v)\n", info, err)
    }
}
