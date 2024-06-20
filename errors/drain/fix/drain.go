package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    ch := make(chan Message)

    // Pupluate some data
    go func() {
        defer close(ch)
        for i := 0; i < 5; i++ {
            msg := Message{
                Type: "test",
                Data: []byte(fmt.Sprintf("payload %d", i)),
            }
            ch <- msg
        }
    }()

    drain(ch, testHandler)
    time.Sleep(time.Second) // let goroutines run
    fmt.Println("DONE")
}

func testHandler(msg Message) {
    ts := msg.Time.Format("15:04:03")
    log.Printf("%s: %s: %x...\n", ts, msg.Type, msg.Data[:20])
}

func drain(ch <-chan Message, handler func(Message)) {
    for msg := range ch {
        msg.Time = time.Now()
        safelyGo(func() {
            handler(msg)
        })
    }
}


type Message struct {
    Time time.Time
    Type string
    Data []byte
}

// safelyGo will run fn in a goroutine, and guard it from panics
func safelyGo(fn func()) {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("error: %s", err)
            }
        }()
        fn()
    }()
}

