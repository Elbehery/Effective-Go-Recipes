package main

import (
    "encoding/gob"
    "fmt"
    "io"
    "log"
    "os"
    "time"
)


// Kind is event kind.
type Kind string

const (
    AddKind      Kind = "add"
    CheckoutKind Kind = "checkout"
)

type Event interface {
    Kind() Kind
}



// Add is an event of adding item to cart.
type Add struct {
    Time time.Time
    ID   string
    User string
    Item int // SKU
}

func (*Add) Kind() Kind {
    return AddKind
}

// Checkout is event for user checkout
type Checkout struct {
    Time time.Time
    Cart string
    User string
}

func (*Checkout) Kind() Kind {
    return CheckoutKind
}



// Encoder is an event encoder.
type Encoder struct {
    enc *gob.Encoder 
}

func NewEncoder(w io.Writer) *Encoder {
    return &Encoder{gob.NewEncoder(w)}
}

func (e *Encoder) Encode(evt Event) error {
    return e.enc.Encode(&evt) 
}


func eventHandler(r io.Reader) error {
    dec := gob.NewDecoder(r)
    for {
        var e Event
        err := dec.Decode(&e)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        switch e.Kind() {
        case AddKind:
            handleAdd(e.(*Add))
        case CheckoutKind:
            handleCheckout(e.(*Checkout))
        default:
            return fmt.Errorf("unknown event kind: %s", e.Kind())
        }
    }

    return nil
}


func handleAdd(a *Add) {
    fmt.Printf("ADD: %#v\n", a)
}

func handleCheckout(c *Checkout) {
    fmt.Printf("CHECKOUT: %#v\n", c)
}

func init() {
    gob.Register(&Add{})
    gob.Register(&Checkout{})
}


func main() {
    r, w, err := os.Pipe()
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    enc := NewEncoder(w)

    go func() {
        defer w.Close()
        a := &Add{time.Now(), "crt42", "joe", 30239}
        if err := enc.Encode(a); err != nil {
            log.Fatalf("error: %s", err)
        }

        c := &Checkout{time.Now(), "crt42", "joe"}
        if err := enc.Encode(c); err != nil {
            log.Fatalf("error: %s", err)
        }
    }()

    if err := eventHandler(r); err != nil {
        log.Fatalf("error: %s", err)
    }
}
