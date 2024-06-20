package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "regexp"
)

var (
    // h96-158.ccnet.com - - [01/Aug/1995:01:36:42 -0400]
    // "GET / HTTP/1.0" 200 7280
    lineRe = regexp.MustCompile(`"(GET|POST) (\S+).*?" (\d+)`)
)


func isError(status string) bool {
    // Either 4xx or 5xx
    return status[0] == '4' || status[0] == '5'
}


func processLog(fileName string, r io.Reader) error {

    s := bufio.NewScanner(r)
    lnum := 0
    for s.Scan() {
        lnum++
        fields := lineRe.FindStringSubmatch(s.Text())
        if fields == nil {
            continue
        }
        url, status := fields[2], fields[3]
        if isError(status) {
            fmt.Printf("%s:%d: %s %s\n", fileName, lnum, url, status)
        }
    }

    if err := s.Err(); err != nil {
        return err
    }
    return nil
}


// ProcessLogFile prints lines with HTTP error status (4xx or 5xx).
func ProcessLogFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    return processLog(path, file)
}


func main() {
    if len(os.Args) < 2 {
        log.Fatalf("usage: %s FILE [FILE ...]", os.Args[0])
    }

    for _, file := range os.Args[1:] {
        if err := ProcessLogFile(file); err != nil {
            log.Fatalf("error: %s", err)
        }
    }
}
