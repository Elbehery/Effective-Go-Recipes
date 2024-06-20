package main

import (
    "encoding/json"
    "fmt"

    "github.com/mitchellh/mapstructure"
)


// Login is a login event.
type Login struct {
    User    int // user ID
    Success bool
}

// Message is a message event.
type Message struct {
    From int // user ID
    To   int // user ID
    Text string
}


func handler(data []byte) error {

    var obj map[string]any
    if err := json.Unmarshal(data, &obj); err != nil {
        return err
    }

    val, ok := obj["type"]
    if !ok {
        return fmt.Errorf("missing `type` in %+v", obj)
    }

    typ, ok := val.(string)
    if !ok {
        return fmt.Errorf("`type` is not a string - %v of %T", val, val)
    }

    switch typ {
    case "login":
        var l Login
        if err := mapstructure.Decode(obj, &l); err != nil {
            return err
        }
        return handleLogin(l)
    case "message":
        var m Message
        if err := mapstructure.Decode(obj, &m); err != nil {
            return err
        }
        return handleMessage(m)
    default:
        return fmt.Errorf("unknown message type: %q", typ)
    }
}

func handleLogin(l Login) error {
    fmt.Printf("login: %#v\n", l)
    return nil
}

func handleMessage(m Message) error {
    fmt.Printf("message: %#v\n", m)
    return nil
}

func main() {
    data := []byte(`{"type": "message", "from": 7, "to": 3, "text": "banana"}`)
    if err := handler(data); err != nil {
        fmt.Println("ERROR:", err)
    }

    data = []byte(`{"type": "login", "user": 7}`)
    if err := handler(data); err != nil {
        fmt.Println("ERROR:", err)
    }

    data = []byte(`{"type": "logout", "user": 7}`)
    if err := handler(data); err != nil {
        fmt.Println("ERROR:", err)
    }
}
