package user


// User is a user message.
type User struct {
    Name string `json:"login"`
    ID   int    `json:"uid"`
}

