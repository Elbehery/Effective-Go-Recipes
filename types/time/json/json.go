package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)


// LogRecord is a log record.
type LogRecord struct {
    Time    time.Time
    Level   string
    Message string
}



// JSONTimeLayout is time format in JSON.
const JSONTimeLayout = "20060102T150405.000"

// JSONTime is a time with different JSON encoding format.
type JSONTime struct {
    time.Time
}

// MarshalJSON implements json.Marshaler
func (t JSONTime) MarshalJSON() ([]byte, error) {
    s := t.Format(JSONTimeLayout)
    return []byte(`"` + s + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *JSONTime) UnmarshalJSON(data []byte) error {
    // Example input: "20230727T170756.246"
    if len(data) < 2 {
        return fmt.Errorf("data too small: %q", string(data))
    }

    data = data[1 : len(data)-1] // trim enclosing ""
    ts, err := time.Parse(JSONTimeLayout, string(data))
    if err != nil {
        return err
    }

    t.Time = ts
    return nil
}



// APILogRecord is a log record used in the API.
type APILogRecord struct {
    Time    JSONTime `json:"time"`
    Level   string   `json:"level"`
    Message string   `json:"message"`
}



// APItoLog converts APILogRecord to LogRecord.
func APItoLog(r APILogRecord) LogRecord {
    lr := LogRecord{
        Time:    r.Time.Time,
        Level:   r.Level,
        Message: r.Message,
    }
    return lr
}

// LogtoAPI converts LogRecord to APILogRecord.
func LogtoAPI(r LogRecord) APILogRecord {
    ar := APILogRecord{
        Time:  JSONTime{r.Time},
        Level: r.Level, Message: r.Message,
    }
    return ar
}


func main() {
    r := APILogRecord{
        Time:    JSONTime{time.Unix(1585318985, 79993962)},
        Level:   "INFO",
        Message: "something happened",
    }
    data, err := json.Marshal(r)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Println(string(data))

    var r2 APILogRecord
    if err := json.Unmarshal(data, &r2); err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Printf("%+v\n", r2)
}
