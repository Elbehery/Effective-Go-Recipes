package main

import (
    "encoding/binary"
    "fmt"
    "log"
    "net"
    "time"
)

var (
    NTPDelta = ntpDelta()
    zeroTime time.Time
)



// NTPMessage is an NTP message.
type NTPMessage struct {
    VNMode           uint8
    Stratum          uint8
    Poll             uint8
    Precision        uint8
    RootDelay        uint32
    RootDispersion   uint32
    RefID            uint32
    RefTimeSec       uint32
    RefTimeFrac      uint32
    OrigTimeSec      uint32
    OrigTimeFrac     uint32
    ReceivedTimeSec  uint32
    ReceivedTimeFrac uint32
    TransmitTimeSec  uint32
    TransmitTimeFrac uint32
}



// TransmitTime returns the transmit time.
func (n *NTPMessage) TransmitTime() time.Time {
    secs := int64(n.TransmitTimeSec)
    nanos := (int64(n.TransmitTimeFrac) * 1e9) >> 32
    return time.Unix(secs, nanos).Add(NTPDelta)
}



// CurrentTime returns the current time from NTP host.
func CurrentTime(addr string) (time.Time, error) {
    conn, err := net.Dial("udp", addr)
    if err != nil {
        return zeroTime, err
    }
    defer conn.Close()

    msg := NTPMessage{
        VNMode: 0b00011011, //  li = 0, vn = 3, and mode = 3
    }

    if err := binary.Write(conn, binary.BigEndian, msg); err != nil {
        return zeroTime, err
    }

    if err := binary.Read(conn, binary.BigEndian, &msg); err != nil {
        return zeroTime, err
    }

    return msg.TransmitTime(), nil
}

func ntpDelta() time.Duration {
    unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
    ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
    return ntpEpoch.Sub(unixEpoch)
}


func main() {
    addr := "pool.ntp.org:123"
    t, err := CurrentTime(addr)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Println(t)
}
