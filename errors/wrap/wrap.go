package main

import (
    "fmt"
    "os"
    "runtime"
    "strings"
)


// Wrapper wraps an error with call stack information.
type Wrapper struct {
    error
    stack []uintptr
}

// Unwrap returns the wrapped error.
func (w *Wrapper) Unwrap() error {
    return w.error
}



// Frame is a call location.
type Frame struct {
    Function string
    File     string
    Line     int
}

var pathSep = string(os.PathSeparator)

func trimPath(path string, size int) string {
    fields := strings.Split(path, pathSep)
    if n := len(fields); n > size {
        fields = fields[len(fields)-size:]
    }
    return strings.Join(fields, pathSep)
}

func (f Frame) String() string {
    return fmt.Sprintf("%s:%d: %s", trimPath(f.File, 3), f.Line, f.Function)
}



// Stack returns the call stack, innermost frame first.
func (w *Wrapper) Stack() []Frame {
    locs := make([]Frame, 0, len(w.stack))
    frames := runtime.CallersFrames(w.stack)
    for {
        frame, more := frames.Next()
        loc := Frame{
            Function: frame.Function,
            File:     frame.File,
            Line:     frame.Line,
        }
        locs = append(locs, loc)

        if !more {
            break
        }
    }
    return locs
}



// Wrap wraps an error with call stack information.
func Wrap(err error) error {
    const depth = 32
    var pcs [depth]uintptr
    n := runtime.Callers(2, pcs[:]) 

    w := Wrapper{
        error: err,
        stack: pcs[:n],
    }
    return &w
}


func a() error {
    return b()
}

func b() error {
    return c()
}

func c() error {
    err := fmt.Errorf("banana")
    return Wrap(err)
}

func main() {
    err := a()
    w := err.(*Wrapper)
    for _, loc := range w.Stack() {
        fmt.Println(loc)
    }
}
