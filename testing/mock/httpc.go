package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)


// APIClient is an API client.
type APIClient struct {
    baseURL string
    c       *http.Client
}

// NewAPIClient returns a new APIClient.
func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{baseURL, &http.Client{}}
}



// Users returns the list of users.
func (c *APIClient) Users() ([]string, error) {
    url := fmt.Sprintf("%s/users", c.baseURL)
    resp, err := c.c.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var users []string
    if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
        return nil, err
    }

    return users, nil
}


func main() {
    c := NewAPIClient("http://localhost:8080")
    users, err := c.Users()
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Println(users)
}
