package main

import (
    "encoding/json"
    "fmt"
    "time"
)


// Log is a log event.
type Log struct {
    Time    time.Time
    Level   int
    Message string
}


func parseLog(data []byte) (*Log, error) {
    var l struct {
        Time    *time.Time
        Level   *int
        Message string
    }
    if err := json.Unmarshal(data, &l); err != nil {
        return nil, err
    }

    if l.Time == nil {
        return nil, fmt.Errorf("missing `time` field")
    }

    if l.Level == nil {
        return nil, fmt.Errorf("missing `level` field")
    }

    if l.Message == "" {
        return nil, fmt.Errorf("missing `message` field")
    }

    return &Log{*l.Time, *l.Level, l.Message}, nil
}


func main() {
    // Valid data
    data := []byte(`{"Time":"2020-10-28T16:58:21.788989027+02:00","Level":20,"Message":"Database is up"}`)
    l, err := parseLog(data)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
    } else {
        fmt.Printf("%#v\n", l)
    }

    // Missing data
    data = []byte(`{"Level":20,"Message":"Database is up"}`)
    l, err = parseLog(data)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
    } else {
        fmt.Printf("%#v\n", l)
    }

}
