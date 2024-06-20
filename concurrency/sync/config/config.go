package main

import (
    "fmt"
    "log"
    "os"
    "sync"
)

var (
    cfgOnce sync.Once
    Config  struct {
        ListenAddress string
        Verbose       bool
    }
)



// LoadConfig loads configuration once from environment.
func LoadConfig(envPrefix string) {
    cfgOnce.Do(func() {
        loadConfig(envPrefix)
    })
}


func loadConfig(envPrefix string) {
    // Utility function to get from environment with prefix
    getEnv := func(name string) string {
        key := fmt.Sprintf("%s_%s", envPrefix, name)
        return os.Getenv(key)

    }

    addr := getEnv("ADDRESS")
    if len(addr) == 0 { // default
        Config.ListenAddress = ":8080"
    } else {
        Config.ListenAddress = addr
    }

    verbose := getEnv("VERBOSE")
    if verbose == "1" || verbose == "yes" || verbose == "on" {
        Config.Verbose = true
    }
    log.Printf("configuration loaded (prefix=%s): %+v", envPrefix, Config)
}


func main() {
    LoadConfig("HTTPD")
    LoadConfig("HTTPD")
}
