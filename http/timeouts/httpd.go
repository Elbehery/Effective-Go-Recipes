package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)


// Post is a posted item.
type Post struct {
    Time    time.Time
    Content string
    Author  string
}


func getLatestPosts(ctx context.Context) ([]Post, error) {
    // FIXME
    return nil, nil
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
    posts, err := getLatestPosts(r.Context())
    if err != nil {
        msg := fmt.Sprintf("can't get posts - %s", err.Error())
        http.Error(w, msg, http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(posts); err != nil {
        log.Printf("error: encode - %s", err)
    }
}


func main() {
    http.HandleFunc("/posts", postsHandler)

    srv := http.Server{
        Addr:         ":8080",
        Handler:      http.DefaultServeMux,
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 2 * time.Second,
    }

    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("error: %s", err)
    }
}

