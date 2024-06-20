package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "math/rand"
    "net/http"
    "time"
)

const (
    jsonMimeType = "application/json"
    csvMimeType  = "text/csv"
)

var (
    // mime type → encoder function
    registry = make(map[string]Encoder)
)

// Encoder function
type Encoder func(w io.Writer, metrics []Metric) error



// Register registers an encoder for a mime type.
// It'll panic if format already registered.
func Register(mimeType string, enc Encoder) error {
    if _, ok := registry[mimeType]; ok {
        return fmt.Errorf("%q already registered", mimeType)
    }
    registry[mimeType] = enc
    return nil
}


// Metric is a measurement of an application behaviour
type Metric struct {
    Time  time.Time
    Name  string
    Value float64
}


// EncodeJSON encodes metrics in JSON format to w.
func EncodeJSON(w io.Writer, metrics []Metric) error {
    return json.NewEncoder(w).Encode(metrics)
}

// EncodeCSV encodes metrics in CSV format to w
func EncodeCSV(w io.Writer, metrics []Metric) error {
    wtr := csv.NewWriter(w)

    // Write header
    if err := wtr.Write([]string{"time", "name", "value"}); err != nil {
        return err
    }
    // Record to write
    r := make([]string, 3)
    for _, s := range metrics {
        r[0] = s.Time.Format(time.RFC3339)
        r[1] = s.Name
        r[2] = fmt.Sprintf("%f", s.Value)
        if err := wtr.Write(r); err != nil {
            return err
        }
    }
    wtr.Flush()
    return nil
}


func init() {
    Register(csvMimeType, EncodeCSV)
    Register(jsonMimeType, EncodeJSON)
}


func queryDB(query string) ([]Metric, error) {
    // FIXME: Currently returns dummy metrics
    var metrics []Metric
    now := time.Now()
    for i := 0; i < 7; i++ {
        m := Metric{
            now.Add(-time.Duration(i) * time.Minute),
            "CPU",
            rand.Float64() * 100,
        }
        metrics = append(metrics, m)
    }
    return metrics, nil
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
    // Get query string from HTTP "query" parameter
    query := r.URL.Query().Get("query")
    if query == "" {
        http.Error(w, "missing query", http.StatusBadRequest)
        return
    }

    mimeType := requestMimeType(r)

    enc, ok := registry[mimeType]
    if !ok {
        msg := fmt.Sprintf("unsupported mime type - %q", mimeType)
        http.Error(w, msg, http.StatusBadRequest)
        return
    }

    metrics, err := queryDB(query)
    if err != nil {
        msg := fmt.Sprintf("can't query - %s", err)
        http.Error(w, msg, http.StatusBadRequest)
        return
    }

    w.Header().Add("Content-Type", mimeType)
    if err := enc(w, metrics); err != nil {
        size := len(metrics)
        const format = "can't encode %d metrics with %q - %s"
        log.Printf(format, size, mimeType, err)
    }
}

func requestMimeType(r *http.Request) string {
    // Get mime type from HTTP Accept header
    mimeType := r.Header.Get("Accept")
    if mimeType == "" || mimeType == "*/*" {
        return jsonMimeType // default to JSON
    }
    return mimeType
}


func main() {
    http.HandleFunc("/metrics", queryHandler)
    addr := ":8080"
    log.Printf("server ready on %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}

