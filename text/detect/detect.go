package main

import (
    "fmt"
    "io"
    "log"
    "mime"
    "net/http"
    "os"

    "golang.org/x/net/html/charset"
)

// ctypeEncoding gets encoding from HTTP Content-Type header
func ctypeEncoding(ctype string) string {
    _, params, err := mime.ParseMediaType(ctype)
    if err != nil {
        return ""
    }
    return params["charset"]
}


func dataEncoding(data []byte) string {
    _, name, certain := charset.DetermineEncoding(data, "text/plain")
    if certain {
        return name
    }
    return ""
}


func main() {
    if len(os.Args) != 2 {
        log.Fatalf("error: %s", "error: wrong number of arguments")
    }

    url := os.Args[1]
    log.Printf("getting %s", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer resp.Body.Close()

    enc := ctypeEncoding(resp.Header.Get("Content-Type"))
    if enc != "" {
        fmt.Printf("Content-Type encoding is %s\n", enc)
        os.Exit(0)
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    enc = dataEncoding(data)
    if enc != "" {
        fmt.Printf("detected encoding is %s\n", enc)
        os.Exit(0)
    }

    fmt.Println("can't detect encoding")
    os.Exit(1)
}
