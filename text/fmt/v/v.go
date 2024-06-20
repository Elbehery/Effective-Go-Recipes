package main

import (
    "fmt"
)

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

    fmt.Printf("%v\n", evt)
    // &{McQueen95 40.7060361 -74.0110143}
    fmt.Printf("%+v\n", evt)
    // &{ID:McQueen95 Latitude:40.7060361 Longitude:-74.0110143}
    fmt.Printf("%#v\n", evt)
    // &main.LocationEvent{ID:"McQueen95", Latitude:40.7060361, \
    //    Longitude:-74.0110143}

    id1, id2 := 1, "1"
    fmt.Printf("id1=%v id2=%v\n", id1, id2)   // id1=1 id2=1
    fmt.Printf("id1=%#v id2=%#v\n", id1, id2) // id1=1 id2="1"
}
