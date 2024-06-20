package main

import (
    "encoding/json"
    "io"
    "log"
    "math/rand"
    "net/http"
    "os"
    "time"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    data, err := io.ReadAll(io.LimitReader(r.Body, maxSize))
    if err != nil {
        http.Error(w, "can't read body", http.StatusBadRequest)
        return
    }
    location := string(data)

    conn, err := Dial(dbDSN)
    if err != nil {
        http.Error(w, "can't connect", http.StatusInternalServerError)
        return
    }
    ch := conn.QueryRidesIn(location)

    rp, wp, err := os.Pipe()
    if err != nil {
        http.Error(w, "can't create pipe", http.StatusInternalServerError)
        return
    }

    go encodeRides(ch, wp)
    _, err = io.Copy(w, rp)
    if err != nil {
        // Can't send error to client here
        log.Printf("error: can't encode: %s", err)
    }
}

// encodeRides encodes rides from ch into w
func encodeRides(ch <-chan Ride, w io.WriteCloser) error {
    enc := json.NewEncoder(w)
    defer w.Close() // Signal no-more-data on function exit

    for r := range ch {
        if err := enc.Encode(r); err != nil {
            return err
        }
    }

    return nil
}



// Ride is an information on a ride.
type Ride struct {
    Time     time.Time
    Distance float64
    Price    float64
}


// Test dummies from here


// QueryRidesIn returns a channel of Rides location.
// It'll close it once there's no more data.
func (c *Conn) QueryRidesIn(location string) <-chan Ride {
    ch := make(chan Ride)

    go func() {
        defer close(ch)
        for i := 0; i < 7; i++ {
            r := Ride{
                Time:     time.Now(),
                Distance: rand.Float64()*10 + 1,
                Price:    rand.Float64()*30 + 5,
            }
            ch <- r
            time.Sleep(time.Second)
        }
    }()

    return ch
}

func Dial(dsn string) (*Conn, error) {
    return &Conn{}, nil
}

// Conn is a dummy connection.
type Conn struct{}

const (
    dbDSN   = "postgresl://localhost:5432"
    maxSize = 1 << 20 // 1MB
)

func main() {
    http.HandleFunc("/query", queryHandler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
