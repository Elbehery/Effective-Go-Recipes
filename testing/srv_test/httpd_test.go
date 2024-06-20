package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
    "os/exec"
    "path"
    "testing"
    "time"
)

func freePort(t *testing.T) int {
    conn, err := net.Listen("tcp", "")
    if err != nil {
        t.Fatal(err)
    }

    conn.Close()
    return conn.Addr().(*net.TCPAddr).Port
}


func waitForServer(t *testing.T, addr string) {
    start := time.Now()
    timeout := 10 * time.Second
    var err error
    var conn net.Conn
    for time.Since(start) < timeout {
        conn, err = net.Dial("tcp", addr)
        if err == nil {
            conn.Close()
            return
        }
        time.Sleep(10 * time.Millisecond)
    }

    t.Fatalf("server not ready after %s (%s)", timeout, err)
}


func buildServer(t *testing.T) string {
    fileName := path.Join(t.TempDir(), "httpd")
    cmd := exec.Command("go", "build", "-o", fileName, "httpd.go")
    if err := cmd.Run(); err != nil {
        t.Fatal(err)
    }

    return fileName
}


func runServer(t *testing.T) int {
    exe := buildServer(t)
    t.Logf("server exe: %q", exe)
    port := freePort(t)
    t.Logf("server port: %d", port)

    env := os.Environ()
    env = append(env, fmt.Sprintf("HTTPD_ADDR=:%d", port))
    cmd := exec.Command(exe)
    cmd.Env = env 
    err := cmd.Start()
    if err != nil {
        t.Fatal(err)
    }

    addr := fmt.Sprintf("localhost:%d", port)
    waitForServer(t, addr)
    t.Cleanup(func() {
        if err := cmd.Process.Kill(); err != nil { 
            t.Logf("warning: can't kill server (pid=%d)", cmd.Process.Pid)
        }
    })

    return port
}


func TestHealth(t *testing.T) {
    port := runServer(t)

    url := fmt.Sprintf("http://localhost:%d/health", port)
    resp, err := http.Get(url)
    if err != nil {
        t.Fatal(err)
    }

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("bad return code: %d", resp.StatusCode)
    }
}

