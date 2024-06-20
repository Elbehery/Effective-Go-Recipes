package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "strings"
)

func kill(pid int) error {
    proc, err := os.FindProcess(pid)
    if err != nil {
        return err
    }

    return proc.Kill()
}


// killServer kills a server from process ID in pidFile
// On success, it'll delete pidFile
func killServer(pidFile string) error {

    file, err := os.Open(pidFile)
    if err != nil {
        return fmt.Errorf("can't open PID file: %w", err)
    }
    defer file.Close()

    var pid int
    _, err = fmt.Fscanf(file, "%d", &pid)
    if err != nil {
        return fmt.Errorf("bad PID in %q: %w", pidFile, err)
    }

    if err := os.Remove(pidFile); err != nil {
        log.Printf("can't remove %q - %s", pidFile, err) // warn, no error
    }

    return kill(pid)
}

func killFromFiles(logFiles []string) error {
    for _, logFile := range logFiles {
        err := killServer(logFile)
        if err == nil {
            return nil
        }

        if !errors.Is(err, os.ErrNotExist) { // File not found
            return err
        }
    }

    files := strings.Join(logFiles, ", ")
    return fmt.Errorf("no existing file found: %v", files)
}

func main() {
    // err := killServer("server.pid")
    logFiles := []string{"server.pid", "httpd.pid"}
    err := killFromFiles(logFiles)
    if err != nil {
        log.Fatalf("%+v", err)
    }
}
