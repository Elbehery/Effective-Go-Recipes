package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/google/uuid"
    "github.com/gorilla/mux"
)


// Message is a channel message.
type Message struct {
    From    string `json:"from"`
    Channel string `json:"channel"`
    Body    string `json:"body"`

    ID string `json:"-"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    var m Message
    if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
        log.Printf("error: bad JSON - %s", err)
        http.Error(w, "bad JSON", http.StatusBadRequest)
    }

    m.ID = uuid.NewString()
    // Adding code redacted

    resp := map[string]any{
        "id":   m.ID,
        "time": time.Now().UTC(),
    }
    data, err := json.Marshal(resp)
    if err != nil {
        log.Printf("error: can't marshal JSON - %s", err)
        http.Error(w, "can't marshal JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    if _, err := w.Write(data); err != nil {
        log.Printf("error: can't send - %s", err)
    }
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/messages", addHandler).Methods("POST")

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("error: %s", err)
    }
}
