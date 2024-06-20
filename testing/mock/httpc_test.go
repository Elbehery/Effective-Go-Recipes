package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/require"
)

type MockTransport struct {
    body []byte
    err  error
}

// RoundTrip implements http.RoundTripper
func (t *MockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
    if t.err != nil {
        return nil, t.err
    }

    w := httptest.NewRecorder() 
    if t.body != nil {
        w.Write(t.body)
    }
    return w.Result(), nil
}


func TestUsersConnectionError(t *testing.T) {
    c := NewAPIClient("http://localhost:8080")
    c.c.Transport = &MockTransport{nil, fmt.Errorf("network error")}
    _, err := c.Users()
    require.Error(t, err)
}


func TestUsersBadJSON(t *testing.T) {
    c := NewAPIClient("http://localhost:8080")
    c.c.Transport = &MockTransport{[]byte(`["clark","diana","bruce"`), nil}
    _, err := c.Users()
    require.Error(t, err)
}


func TestUsersOK(t *testing.T) {
    users := []string{"clark", "diana", "bruce"}
    data, err := json.Marshal(users)
    require.NoError(t, err, "encode JSON")

    c := NewAPIClient("http://localhost:8080")
    c.c.Transport = &MockTransport{data, nil}
    reply, err := c.Users()
    require.NoError(t, err, "API call")
    require.Equal(t, users, reply, "users")
}

