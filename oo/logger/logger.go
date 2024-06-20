package main

import (
    "fmt"
    "io"
    "os"
    "time"
)


// Level is a log level.
type Level byte

const (
    Debug Level = iota + 1
    Info
    Error
)

func (l Level) String() string {
    switch l {
    case Debug:
        return "DEBUG"
    case Info:
        return "INFO"
    case Error:
        return "ERROR"
    }

    return fmt.Sprintf("Level <%d>", l)
}

type Logger struct {
    level Level
    w     io.Writer
}

func NewLogger(level Level, out io.Writer) Logger {
    l := Logger{level, out}
    return l
}


func (l Logger) log(level Level, format string, args ...any) {
    if level < l.level {
        return
    }

    msg := fmt.Sprintf(format, args...)
    ts := time.Now().UTC().Format(time.RFC3339)
    fmt.Fprintf(l.w, "[%s] - %s - %s\n", ts, level, msg)
}



// Debug is a debug log.
func (l Logger) Debug(format string, args ...any) {
    l.log(Debug, format, args...)
}

// Info is an info log.
func (l Logger) Info(format string, args ...any) {
    l.log(Info, format, args...)
}

// Error is an error log.
func (l Logger) Error(format string, args ...any) {
    l.log(Error, format, args...)
}


func main() {
    log := NewLogger(Info, os.Stdout)
    log.Debug("debug")
    log.Info("info")
    log.Error("error")
}
