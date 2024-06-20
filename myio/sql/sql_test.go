package sql

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenSelect(t *testing.T) {
	testCases := []struct {
		name    string
		table   string
		columns []string
		exp     string
		expErr  error
	}{
		{
			name:    "correct input",
			table:   "table",
			columns: []string{"A", "B", "C"},
			exp:     "SELECT\n\tA,\n\tB,\n\tC\nFROM table",
			expErr:  nil,
		},
		{
			name:    "no columns",
			table:   "table",
			columns: []string{},
			exp:     "",
			expErr:  fmt.Errorf("empty columns"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := genSelect(tc.table, tc.columns)
			require.Equal(t, tc.expErr, err)
			require.Equal(t, tc.exp, act)
		})
	}
}
