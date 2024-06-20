package main

import (
    "fmt"
    "log"
    "time"
)


// StockInfo is information about stock at a given date.
type StockInfo struct {
    Date   time.Time
    Symbol string
    Open   float64
    High   float64
    Low    float64
    Close  float64
}


type key struct {
    year   int
    month  time.Month
    day    int
    symbol string
}



// InfoDB is in memory stock database.
type InfoDB struct {
    m map[key]StockInfo
}



// NewInfoDB return new InfoDB
func NewInfoDB(info []StockInfo) (*InfoDB, error) {
    m := make(map[key]StockInfo)
    for _, si := range info {
        k := key{si.Date.Year(), si.Date.Month(), si.Date.Day(), si.Symbol}
        if _, ok := m[k]; ok {
            return nil, fmt.Errorf("duplicate info: %#v", si)
        }
        m[k] = si
    }

    return &InfoDB{m}, nil
}



// Get return information for stock at date. If information is not found the
// second return value will be false.
func (i *InfoDB) Get(symbol string, date time.Time) (StockInfo, bool) {
    k := key{date.Year(), date.Month(), date.Day(), symbol}
    info, ok := i.m[k]
    return info, ok
}


func day(year int, month time.Month, day int) time.Time {
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func main() {
    info := []StockInfo{
        {day(2020, time.March, 2), "AAPL", 107.32, 107.05, 108.5, 108.2},
        {day(2020, time.March, 3), "AAPL", 108.32, 108.05, 109.5, 109.2},
    }

    db, err := NewInfoDB(info)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    d := day(2020, time.March, 2)
    fmt.Println(db.Get("AAPL", d))

    d = day(2020, time.March, 20)
    fmt.Println(db.Get("AAPL", d))

}
