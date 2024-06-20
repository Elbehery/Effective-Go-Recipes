package sql

import (
	"bytes"
	"fmt"
)

func genSelect(table string, columns []string) (string, error) {
	var buf bytes.Buffer

	if len(columns) == 0 {
		return "", fmt.Errorf("empty columns")
	}

	fmt.Fprintln(&buf, "SELECT")
	sep := ","
	for idx, col := range columns {
		if idx == len(columns)-1 {
			sep = ""
		}
		fmt.Fprintf(&buf, "	%s%s\n", col, sep)
	}

	fmt.Fprintf(&buf, "FROM %s", table)
	return buf.String(), nil
}
