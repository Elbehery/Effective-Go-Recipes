package main

import (
    "os"
    "testing"
)

var (
    inCi = os.Getenv("CI") != ""
)


func TestMonthlyReport(t *testing.T) {
    if !inCi {
        t.Skip("not in CI system")
    }

    t.Log("running in CI")
    // TODO: Rest of long running test ...
}

