package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/rs/xid"
)

type ctxKey string

const (
    valuesKey ctxKey = "ctxArgs"
)

type Values struct {
    RequestID string
    Logger    *log.Logger
}


func idLogger(id string) *log.Logger {
    prefix := fmt.Sprintf("<%s> %s", id, log.Prefix())
    return log.New(log.Writer(), prefix, log.Flags())
}


func ctxLogger(ctx context.Context) *log.Logger {
    vals, ok := ctx.Value(valuesKey).(*Values) 
    if !ok {
        return stdLogger()
    }

    if vals.Logger == nil {
        panic(fmt.Sprintf("no logger in %#v", vals))
    }
    return vals.Logger
}

// stdLogger returns a logger that behaves like the top-level function in
// the log package
func stdLogger() *log.Logger {
    return log.New(log.Writer(), log.Prefix(), log.Flags())
}


func newID() string {
    return xid.New().String()
}


func usersHandler(w http.ResponseWriter, r *http.Request) {
    timeout := 100 * time.Millisecond
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    id := newID()
    logger := idLogger(id)
    logger.Printf("info: usersHandler: id=%s", id)

    values := Values{
        RequestID: id,
        Logger:    logger,
    }
    ctx = context.WithValue(ctx, valuesKey, &values)
    users, err := getAllUsers(ctx)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    logger.Printf("info: usersHandler: got %d users", len(users))

    json.NewEncoder(w).Encode(users)
}

func getAllUsers(ctx context.Context) ([]string, error) {
    logger := ctxLogger(ctx)
    logger.Printf("info: getting all users")

    // FIXME: Connect to database, query ...
    return nil, nil
}

func main() {
    http.HandleFunc("/users", usersHandler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("error: %s", err)
    }
}
