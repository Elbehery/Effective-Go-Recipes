package main

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "os/exec"
    "runtime"
    "sort"
    "strconv"
)

func median(values []float64) float64 {
    // Don't mutate original values
    nums := make([]float64, len(values))
    copy(nums, values)
    sort.Float64s(nums)

    i := len(nums) / 2
    if len(nums)%2 == 0 {
        return (nums[i-1] + nums[i]) / 2.0
    }

    return nums[i]
}


// findTime finds the ping time in ping output line
// returns value, found and error
func findTime(line []byte) (float64, bool, error) {
    // 64 bytes from ...: icmp_seq=1 ttl=51 time=142 ms
    var prefix = []byte("time=")
    start := bytes.Index(line, prefix)
    if start == -1 {
        return 0, false, nil
    }
    start += len(prefix) // skip over "time="
    end := bytes.IndexByte(line[start:], ' ')
    if end == -1 {
        return 0, false, fmt.Errorf("can't find end")
    }

    end += start
    val, err := strconv.ParseFloat(string(line[start:end]), 64)
    if err != nil {
        return 0, false, err
    }

    return val, true, nil
}


// medianPing returns the median time of <count> pings to <host>.
func medianPingTime(host string, count int) (float64, error) {
    sw := "-c"
    if runtime.GOOS == "windows" {
        sw = "-n" // windows ping uses -n for count
    }

    // ping -c 4 pragprog.com
    cmd := exec.Command("ping", sw, fmt.Sprintf("%d", count), host)

    data, err := cmd.Output() 
    if err != nil {           // Wait for ping to finish
        return 0, err
    }

    values := make([]float64, 0, count)
    s := bufio.NewScanner(bytes.NewReader(data)) 
    for s.Scan() {
        val, found, err := findTime(s.Bytes())
        if err != nil {
            return 0, err
        }
        if !found {
            continue
        }
        values = append(values, val)
    }

    if err := s.Err(); err != nil {
        return 0, err
    }

    return median(values), nil
}


func main() {
    mp, err := medianPingTime("pragprog.com", 5)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    fmt.Println(mp)
}
