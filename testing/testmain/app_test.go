package main

import (
    "os"
    "testing"
)

func TestApp(t *testing.T) {
    t.Logf("TODO (%q)", os.Getenv(envConfigKey))
}
