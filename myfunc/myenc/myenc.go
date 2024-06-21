package myenc

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

const (
	jsonMimeType = "application/json"
	csvMimeType  = "text/csv"
)

var (
	registry = make(map[string]Encoder)
)

func init() {
	Register(jsonMimeType, EncodeJson)
	Register(csvMimeType, EncodeCSV)
}

type Metric struct {
	Time time.Time
	Name string
	Val  float64
}

type Encoder func(w io.Writer, metrics []Metric) error

func Register(mimeType string, enc Encoder) error {
	if _, ok := registry[mimeType]; ok {
		return fmt.Errorf("type is already registered")
	}
	registry[mimeType] = enc
	return nil
}

func EncodeJson(w io.Writer, metrics []Metric) error {
	return json.NewEncoder(w).Encode(metrics)
}

func EncodeCSV(w io.Writer, metrics []Metric) error {
	wr := csv.NewWriter(w)

	if err := wr.Write([]string{"Time", "Name", "Val"}); err != nil {
		return err
	}

	rec := make([]string, 3)
	for _, r := range metrics {
		rec[0] = r.Time.Format(time.RFC3339)
		rec[1] = r.Name
		rec[2] = fmt.Sprintf("%f", r.Val)

		if err := wr.Write(rec); err != nil {
			return err
		}
	}
	wr.Flush()
	return nil
}
