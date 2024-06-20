package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "strings"
)

func sendFile(addr, fileName string) error {
    c, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }
    defer c.Close()

    file, err := os.Open(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return err
    }

    // Send header line with file size & name
    _, err = fmt.Fprintf(c, "%s %d ", fileName, info.Size())
    if err != nil {
        return err
    }

    // Send file content
    _, err = io.Copy(c, file)
    if err != nil {
        return err
    }

    const maxReply = 1 << 10 // 1KB
    data, err := io.ReadAll(io.LimitReader(c, maxReply))
    if err != nil {
        return err
    }

    reply := string(data)
    log.Printf("reply: %s", reply)
    if strings.HasPrefix(reply, "error") {
        return fmt.Errorf(reply)
    }

    return nil
}

func main() {
    if err := sendFile("localhost:8765", "client.go"); err != nil {
        log.Fatalf("error: %s", err)
    }

}

