package main

import (
    "fmt"
)


// Priority is bug priority.
type Priority uint8

const (
    Low    Priority = 10
    Medium Priority = 20
    High   Priority = 30
)



// Bug is a bug in the system.
type Bug struct {
    Title    string
    Priority Priority
}



// String implements the fmt.Stringer interface.
func (p Priority) String() string {
    switch p {
    case Low:
        return "low"
    case Medium:
        return "medium"
    case High:
        return "high"
    }

    return fmt.Sprintf("<%d>", p)
}


func main() {
    bug := Bug{"Bug level is printed as number", Medium}
    fmt.Printf("%+v\n", bug)
    // {Title:Bug level is printed as number Priority:medium}
}

