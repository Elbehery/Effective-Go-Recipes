package main

import (
    "sync"
    "time"
)

var (
    cfgLock sync.RWMutex
    config  = make(map[string]string)
)



// ReloadConfig reloads configuration.
func ReloadConfig() {
    cfgLock.Lock()
    defer cfgLock.Unlock()

    config["updated"] = time.Now().String()
    // TODO: finish loading configuration
}



// GetConfig returns the value for key in configuration.
func GetConfig(key string) string {
    cfgLock.RLock()
    defer cfgLock.RUnlock()

    return config[key]
}

