package main

import (
    "fmt"
    "strings"
    "time"

    "golang.org/x/text/unicode/norm"
)


// Tour is a tour in a city.
type Tour struct {
    City string
    Name string
    Time time.Time
}


// normString normalizes string in NFKC format
func normString(s string) string {
    return norm.NFKC.String(s)
}



// NewTour returns a new Tour.
func NewTour(city, name string, time time.Time) *Tour {
    tour := Tour{
        City: normString(city),
        Name: name,
        Time: time,
    }
    return &tour
}


// findTours returns all tours in city
func findTours(db []*Tour, city string) []*Tour {
    city = normString(city) 
    var tours []*Tour
    for _, t := range db {
        if strings.EqualFold(t.City, city) {
            tours = append(tours, t)
        }
    }
    return tours
}


// date is a shortcut to create time.Time from year, month, day
func date(year int, month time.Month, day int) time.Time {
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func main() {
    db := []*Tour{
        {"Gdańsk", "Polish Food", date(2021, 1, 1)},
        {"Kraków", "Pub to Pub", date(2021, 1, 2)},
    }
    tours := findTours(db, "Kraków")
    fmt.Printf("number of tours found: %d\n", len(tours))

    city1, city2 := "Kraków", "Kraków"
    fmt.Println("city1:", len(city1))
    fmt.Println("city2:", len(city2))
}
