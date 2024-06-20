package main

import (
    "sync"
    "sync/atomic"
    "testing"
)

var (
    count int64
    m     sync.Mutex
)

func BenchmarkMutex(b *testing.B) {
    for i := 0; i < b.N; i++ {
        m.Lock()
        count++
        m.Unlock()
    }
}

func BenchmarkAtomic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        atomic.AddInt64(&count, 1)
    }
}
