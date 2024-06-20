package main

import (
    "log"
    "os"
    "testing"

    "gopkg.in/yaml.v2"
)

const (
    envConfigKey = "APP_CONFIG_FILE"
)

func setupTests() error {
    file, err := os.CreateTemp("", "*.yml")
    if err != nil {
        return err
    }
    defer file.Close()

    cfg := map[string]any{
        "verbose": true,
        "dsn":     "postgres://localhost:5432",
    }
    if err := yaml.NewEncoder(file).Encode(cfg); err != nil {
        return err
    }

    os.Setenv(envConfigKey, file.Name())
    log.Printf("config file: %q", file.Name())
    return nil
}


func teardownTests() {
    fileName := os.Getenv(envConfigKey)
    if err := os.Remove(fileName); err != nil {
        log.Printf("warning: can't delete %q - %s", fileName, err)
    }
}


func runTests(m *testing.M) int {
    if err := setupTests(); err != nil {
        log.Printf("error: can't setup tests - %s", err)
        return 1
    }
    defer teardownTests()
    return m.Run()
}


func TestMain(m *testing.M) {
    code := runTests(m)
    os.Exit(code)
}

