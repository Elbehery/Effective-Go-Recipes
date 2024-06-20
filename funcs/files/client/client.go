package client

import "log"

var (
    maxBodySize int64 = 1_000_000
)

func setMaxBodySize(size int64) int64 {
    old := maxBodySize
    log.Printf("maxBodySize %d -> %d", old, size)
    maxBodySize = size
    return old
}

type Client struct{}

func New() *Client {
    return &Client{}
}
