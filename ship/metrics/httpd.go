package main

import (
    "expvar"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type statusWriter struct {
    http.ResponseWriter 
    statusCode          int
}

func (s *statusWriter) WriteHeader(status int) {
    s.statusCode = status
    s.ResponseWriter.WriteHeader(status)
}


func addMetrics(name string, h http.Handler) http.Handler {
    calls := expvar.NewInt(fmt.Sprintf("%s.calls", name))
    errors := expvar.NewInt(fmt.Sprintf("%s.errors", name))
    oks := expvar.NewInt(fmt.Sprintf("%s.oks", name))

    fn := func(w http.ResponseWriter, r *http.Request) {
        calls.Add(1)

        sw := statusWriter{w, http.StatusOK}
        h.ServeHTTP(&sw, r)

        if sw.statusCode >= http.StatusBadRequest {
            errors.Add(1)
        } else {
            oks.Add(1)
        }
    }

    return http.HandlerFunc(fn)
}


func lookupHandler(w http.ResponseWriter, r *http.Request) {
    lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
    if err != nil || lat < -90 || lat > 90 {
        http.Error(w, "bad lat", http.StatusBadRequest)
        return
    }


    lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
    if err != nil || lng < -180 || lng > 180 {
        http.Error(w, "bad lng", http.StatusBadRequest)
        return
    }

    // Lookup code redacted
    fmt.Fprintln(w, "Earth")
}


func main() {
    h := addMetrics("lookup", http.HandlerFunc(lookupHandler))
    http.Handle("/lookup", h)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
