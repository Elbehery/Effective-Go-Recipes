package main

import (
    "fmt"
    "io"
    "os"
    "time"
)


// LogLevel is logging level.
type LogLevel uint8

// Log levels
const (
    DebugLevel LogLevel = iota + 1
    InfoLevel
    ErrorLevel
)

func (l LogLevel) String() string {
    switch l {
    case DebugLevel:
        return "DEBUG"
    case InfoLevel:
        return "INFO"
    case ErrorLevel:
        return "ERROR"
    }

    return fmt.Sprintf("<LogLevel: %d>", l)
}



// Logger is a logger to a file with a given level.
type Logger struct {
    file  string
    w     io.WriteCloser
    level LogLevel
}

func (l *Logger) String() string {
    return fmt.Sprintf("Logger: %s -> %s", l.level, l.file)
}

// Debug emits a debug message
func (l *Logger) Debug(format string, args ...any) {
    l.emit(DebugLevel, format, args...)
}

// Info emits an info message
func (l *Logger) Info(format string, args ...any) {
    l.emit(InfoLevel, format, args...)
}

// Error emits an error message
func (l *Logger) Error(format string, args ...any) {
    l.emit(ErrorLevel, format, args...)
}

func (l *Logger) emit(level LogLevel, format string, args ...any) {
    if level > l.level {
        return
    }
    ts := time.Now().Format(time.RFC3339)
    fmt.Fprintf(l.w, "%s - %s - ", ts, level)
    fmt.Fprintf(l.w, format, args...)
    fmt.Fprint(l.w, "\n")
}



// NewLogger returns a new logger that will log to fileName logs with level
// higher or equal to level.
func NewLogger(fileName string, level LogLevel) (*Logger, error) {
    flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
    file, err := os.OpenFile(fileName, flag, 0644)
    if err != nil {
        return nil, err
    }

    return &Logger{fileName, file, level}, nil
}


func main() {
    log, err := NewLogger("/tmp/a.log", InfoLevel)
    if err != nil {
        panic(err)
    }

    log.Debug("debug message #%d", 1)
    log.Info("info message #%d", 2)
    log.Error("error message #%d", 3)
}
