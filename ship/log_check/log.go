package main

import (
    "fmt"
    "log"
    "time"

    "go.uber.org/zap"
)

func main() {
    cfg := zap.NewProductionConfig()
    cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
    // cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
    logger, err := cfg.Build()
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    s := Server{logger: logger}
    uid, err := s.login(token)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Println(uid)
}

// Login authenticates user from token and returns user id
func (s *Server) login(token string) (int, error) {
    uid, err := auth.Authenticate(token) // API call to auth service
    if err != nil {
        return 0, err
    }

    if info := s.logger.Check(zap.InfoLevel, "logged in"); info != nil {
        info.Write(
            zap.Any("user", users.Load(uid)),
        )
    }

    return uid, nil
}


type Server struct {
    logger *zap.Logger
}

// Below methods are stubs to help the code compile

const token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSmFtZXMgQm9uZCIsImlkIjo3fQ.jrwaqo2qG2-8e4GOZSfGt6sPj-yqqm2HQsEqjr92-Fs`

var auth Auth

func (a *Auth) Authenticate(token string) (int, error) { return 7, nil }

type Auth struct{}

var users Users

type Users struct{}

func (u *Users) Load(id int) User {
    time.Sleep(10 * time.Millisecond)
    return User{}
}

type User struct{}
