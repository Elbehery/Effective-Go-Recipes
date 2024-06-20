package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    "github.com/ardanlabs/conf/v3"
)

func validateAddr(addr string) error {
    i := strings.Index(addr, ":")
    if i == -1 {
        return fmt.Errorf("%q: missing : in address", addr)
    }

    port, err := strconv.Atoi(addr[i+1:])
    if err != nil {
        return fmt.Errorf("%q: invalid port - %w", addr, err)
    }

    const maxPort = 65_535
    if port < 0 || port > maxPort {
        return fmt.Errorf("%q: invalid port number", addr)
    }

    return nil
}


func main() {
    var cfg struct {
        Web struct {
            Addr string `conf:"default::8080,env:ADDR"` 
        }
    }

    help, err := conf.Parse("APP", &cfg)
    if err != nil {
        if errors.Is(err, conf.ErrHelpWanted) {
            fmt.Println(help)
            os.Exit(0)
        }
        log.Fatalf("error: bad config - %s", err)
    }


    if err := validateAddr(cfg.Web.Addr); err != nil {
        log.Fatalf("error: invalid config - %s", err)
    }

    fmt.Printf("%+v\n", cfg)
}
