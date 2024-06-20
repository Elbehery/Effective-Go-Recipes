package main

import (
    "bytes"
    "fmt"
    "log"
)

func genSelect(table string, columns []string) (string, error) {
    var buf bytes.Buffer

    if len(columns) == 0 {
        return "", fmt.Errorf("empty select")
    }

    fmt.Fprintln(&buf, "SELECT")
    for i, col := range columns {
        suffix := ","
        if i == len(columns)-1 {
            suffix = "" // No trailing comma in SQL
        }
        fmt.Fprintf(&buf, "    %s%s\n", col, suffix)
    }

    fmt.Fprintf(&buf, "FROM %s;", table)

    return buf.String(), nil
}


func main() {
    table := "rides"
    columns := []string{"id", "time", "duration", "price"}

    sql, err := genSelect(table, columns)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    fmt.Println(sql)
}
