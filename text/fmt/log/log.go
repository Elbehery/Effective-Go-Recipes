package main

import (
    "log"
)


// LocationEvent is a location event.
type LocationEvent struct {
    ID        string
    Latitude  float64
    Longitude float64
}


func main() {
    evt := &LocationEvent{
        ID:        "McQueen95",
        Latitude:  40.7060361,
        Longitude: -74.0110143,
    }

    log.Printf("info: loc: %#v", evt)
}
