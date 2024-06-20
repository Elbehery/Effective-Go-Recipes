package httpd

import (
    "sync/atomic"
    "time"
)

var (
    now atomic.Value
)



// Now return the current time in 1ms granularity
func Now() time.Time {
    return now.Load().(time.Time)
}


func init() {
    now.Store(time.Now())
    go func() {
        for {
            time.Sleep(time.Millisecond)
            now.Store(time.Now())
        }
    }()
}

