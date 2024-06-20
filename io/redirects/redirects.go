package main

import (
    "bufio"
    "compress/gzip"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// numRedirects returns the total number of lines and the number of lines
// with an HTTP redirect
func numRedirects(r io.Reader) (int, int, error) {
    s := bufio.NewScanner(r)
    nLines, nRedirects := 0, 0
    for s.Scan() {
        nLines++
        // Example:
        // 203.252.212.44 - - [01/Aug/1995:03:45:47 -0400] \
        // "GET /ksc.html HTTP/1.0" 200 7280
        fields := strings.Fields(s.Text())
        code := fields[len(fields)-2] // code is one before last
        if code[0] == '3' {           // HTTP redirect is 3XX
            nRedirects++
        }
    }

    if err := s.Err(); err != nil {
        return -1, -1, err
    }

    return nLines, nRedirects, nil
}


func main() {
    matches, err := filepath.Glob("logs/http-*.log*")
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    nLines, nRedirects := 0, 0
    for _, fileName := range matches {
        file, err := os.Open(fileName)
        if err != nil {
            log.Fatalf("error: %s", err)
        }

        var r io.Reader = file
        if strings.HasSuffix(fileName, ".gz") {
            r, err = gzip.NewReader(r) 
            if err != nil {
                log.Fatalf("%q - %v", fileName, err)
            }
        }

        nl, nr, err := numRedirects(r)
        if err != nil {
            log.Fatalf("%q - %v", fileName, err)
        }
        nLines += nl
        nRedirects += nr
        file.Close()
    }

    fmt.Printf("%d redirects in %d lines\n", nRedirects, nLines)
}
