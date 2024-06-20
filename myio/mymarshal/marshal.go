package mymarshal

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

// Ride is a ride record
type Ride struct {
	Id       int
	Time     time.Time
	Duration time.Duration
	Price    float64
	Distance float64
}

func UnmarshalRide(data []byte, ride *Ride) error {
	r := bytes.NewReader(data)
	return NewDecoder(r).DecodeRide(ride)
}

// Decoder is an example decoder
type Decoder struct {
	dec *json.Decoder
}

// NewDecoder returns a new decoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{dec: json.NewDecoder(r)}
}

func (d *Decoder) DecodeRide(r *Ride) error {
	return d.dec.Decode(r)
}
