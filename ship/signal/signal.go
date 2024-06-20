package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "golang.org/x/sys/unix"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Real health check
    fmt.Fprintln(w, "OK")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)

    addr := ":8080"
    srv := &http.Server{
        Addr:    addr,
        Handler: mux,
    }

    go func() {
        log.Printf("server starting on %s", addr)
        err := srv.ListenAndServe()
        
        
        
        if err != nil && err != http.ErrServerClosed {
            log.Printf("error: listen - %s", err)
            os.Exit(1)
        }
    }()

    ch := make(chan os.Signal, 1)                
    signal.Notify(ch, unix.SIGTERM, unix.SIGINT) 

    <-ch
    log.Printf("shutting down")
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("error: shutdown - %s", err)
    }

}
