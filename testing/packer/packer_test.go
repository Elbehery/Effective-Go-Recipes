package packer

import (
    "fmt"
    "os"
    "testing"

    "gopkg.in/yaml.v2"
)

type TestCase struct {
    Name      string    `yaml:"name"`
    BoxCap    float64   `yaml:"box_capacity"`
    N         int       `yaml:"num_boxes"`
    Weights   []float64 `yaml:"weights"`
    ShouldErr bool      `yaml:"err"`
}


func weightsToItems(ws []float64) []Item {
    items := make([]Item, 0, len(ws))
    for i, w := range ws {
        item := Item{
            Name:   fmt.Sprintf("item %d", i+1),
            Weight: w,
        }
        items = append(items, item)
    }
    return items
}


func loadCases(t *testing.T, path string) []TestCase {
    file, err := os.Open(path)
    if err != nil {
        t.Fatal(err)
    }
    defer file.Close()

    var tcs []TestCase
    if err := yaml.NewDecoder(file).Decode(&tcs); err != nil {
        t.Fatal(err)
    }
    return tcs
}


func TestPack(t *testing.T) {
    testCases := loadCases(t, "packer_cases.yml")
    for _, tc := range testCases {
        items := weightsToItems(tc.Weights)
        t.Run(tc.Name, func(t *testing.T) { 
            boxes, err := Pack(tc.BoxCap, items)
            if err == nil && tc.ShouldErr {
                t.Fatal("expected error, got nil")
            }
            if err != nil && !tc.ShouldErr {
                t.Fatalf("unexpected error: %s", err)
            }

            if n := len(boxes); n != tc.N {
                t.Fatalf("expected %d boxes, got %d", tc.N, n)
            }
        })
    }
}

