package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "path"
    "strings"
)

const (
    maxSize = 10 * (1 << 30) // 10GB
)

func fileHandler(c net.Conn) {
    defer c.Close()

    var fileName string
    var size int64
    if _, err := fmt.Fscanf(c, "%s %d", &fileName, &size); err != nil {
        log.Printf("can't read file size: %s", err)
        fmt.Fprintf(c, "error: %s\n", err)
        return
    }
    log.Printf("%s: %d", fileName, size)

    if size > maxSize {
        log.Printf("size %d > %d", size, maxSize)
        fmt.Fprintf(c, "error: size %d > %d\n", size, maxSize)
        return
    }

    // Save in "logs" directory
    fileName = path.Join("logs", strings.TrimSpace(fileName))
    log.Printf("fileName: %s", fileName)

    file, err := os.Create(fileName)
    if err != nil {
        log.Printf("can't open %s - %s", fileName, err)
        fmt.Fprintf(c, "error: %s\n", err)
        return
    }
    defer file.Close()

    n, err := io.CopyN(file, c, size)
    if err != nil {
        log.Printf("can't copy to %s: %s", fileName, err)
        fmt.Fprintf(c, "error: %s\n", err)
        return
    }

    if n != size {
        log.Printf("got %d bytes of %d", n, size)
        fmt.Fprintf(c, "error: %d of %d bytes", n, size)
        return
    }

    fmt.Fprintf(c, "ok: %d bytes written to %s\n", n, fileName)
}

func main() {
    addr := ":8765"
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    log.Printf("server ready on %s", addr)

    for {
        c, err := ln.Accept()
        if err != nil {
            log.Fatalf("error: %s", err)
        }
        log.Printf("new connection from %s", c.RemoteAddr())
        go fileHandler(c)
    }
}
