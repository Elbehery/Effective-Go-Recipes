package main

import (
    "encoding/json"
    "io"
    "log"
    "net"
    "os"
    "os/exec"
    "time"
)

const (
    socketFile = "/tmp/srv.sock"
)

type Message struct {
    Text string `json:"text"`
}

func processMessages(conn io.ReadWriteCloser, ch <-chan Message) error {

    dec := json.NewDecoder(conn)
    enc := json.NewEncoder(conn)

    for msg := range ch {

        // json encoder add newline
        if err := enc.Encode(msg); err != nil {
            return err
        }

        var reply struct {
            Output any
        }
        if err := dec.Decode(&reply); err != nil {
            return err
        }

        log.Printf("%#v -> %#v", msg, reply.Output)
    }

    return nil
}

func main() {
    cmd := exec.Command(
        "go", "run", "server/main.go",
        "-socket", socketFile,
    )
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Start(); err != nil {
        log.Fatalf("error: %s", err)
    }
    defer cmd.Process.Kill()

    time.Sleep(time.Second) // Give server time to load (FIXME: timeout)

    sock, err := net.Dial("unix", socketFile)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer sock.Close()

    ch := make(chan Message)
    go func() {
        ch <- Message{Text: "Was it a cat I saw?"}
        close(ch)
    }()

    if err := processMessages(sock, ch); err != nil {
        log.Fatalf("error: %s", err)
    }
}
