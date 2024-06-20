package main

import (
    "encoding/json"
    "fmt"
    "log"
)


// Stack is a LIFO data structure.
type Stack struct {
    Value string
    Next  *Stack
}



// MarshalJSON implements json.Marshaler.
func (s *Stack) MarshalJSON() ([]byte, error) {
    var values []string
    for s != nil {
        values = append(values, s.Value)
        s = s.Next
    }

    return json.Marshal(values) 
}



// UnmarshalJSON implements json.Unmarshaler.
func (s *Stack) UnmarshalJSON(data []byte) error {
    var values []string
    if err := json.Unmarshal(data, &values); err != nil {
        return err
    }

    var node *Stack
    for i := len(values) - 1; i >= 0; i-- {
        node = &Stack{values[i], node}
    }

    *s = *node 
    return nil
}


func main() {
    s := &Stack{"Hello", &Stack{"there", nil}}
    data, err := json.Marshal(s)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Println(string(data))

    var s1 *Stack
    if err := json.Unmarshal(data, &s1); err != nil {
        log.Fatalf("error: %s", err)
    }
    for s1 != nil {
        fmt.Println(s1.Value)
        s1 = s1.Next
    }
}
