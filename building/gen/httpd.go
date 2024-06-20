package main

//go:generate go run _scripts/gen_ips.go _scripts/allowed_ips.txt ips.go
//go:generate go fmt ips.go

import (
    "log"
    "net/http"
    "strings"
)

func requestIP(r *http.Request) string {
    fields := strings.Split(r.RemoteAddr, ":")
    if len(fields) != 2 {
        return ""
    }

    return fields[0]
}


func handler(w http.ResponseWriter, r *http.Request) {
    if ip := requestIP(r); !IPAllowed(ip) {
        log.Printf("WARNING: (sec) access from disallowed IP: %s", ip)
        http.Error(w, "Unknown IP", http.StatusUnauthorized)
        return
    }

    w.Write([]byte("You're in!\n"))
}

func main() {
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
