package main

import (
    "encoding/json"
    "fmt"
)


// UserRequest is a request for user information.
type UserRequest struct {
    Login string
}

// GroupRequest is a request for group information.
type GroupRequest struct {
    ID string
}

// Request is the set of all possible requests.
type Request interface {
    UserRequest | GroupRequest
}



// UnmarshalJSON implements json.Unmarshaler.
func UnmarshalJSON[T Request](data []byte, d *T) error { 
    return json.Unmarshal(data, d)
}


func main() {
    data := []byte(`{"login": "elliot"}`)
    var r UserRequest
    fmt.Println(UnmarshalJSON(data, &r)) // nil
    fmt.Println(r)                       // {elliot}

    // fmt.Println(UnmarshalJSON(data, r)) // won't compile
}
