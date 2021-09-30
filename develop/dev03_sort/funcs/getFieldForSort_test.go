package funcs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFieldForSortFromLines(t *testing.T) {
	tables := []struct {
		numberColumn      int
		line              string
		resultSortedField string
		resultSearched    bool
	}{
		{0, "a b c d", "a", true},
		{1, "a b c d", "b", true},
		{2, "a b c d", "c", true},
		{3, "a b c d", "d", true},
		{4, "a b c d", "", false},
		{5, "a b c d", "", false},
		{3, "", "", false},
		{-1, "", "", false},
		{-1, "a b c d", "a b c d", true},
		{2, "a   b    c      d", "c", true},
	}
	for _, table := range tables {
		sortedField, searched := getFieldForSortFromLines(table.numberColumn, table.line)

		inputDataInfo := fmt.Sprintf("number of column = %v, line = (%v)\n", table.numberColumn, table.line)
		assert.Equal(t, sortedField, table.resultSortedField, inputDataInfo)
		assert.Equal(t, searched, table.resultSearched, inputDataInfo)
	}
}
