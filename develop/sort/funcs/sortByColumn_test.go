package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortByColumn(t *testing.T) {
	tables := []struct {
		numberColumn    int
		lines           []string
		neededToReverse bool
		res             []string
	}{
		{-1, []string{"b", "c", "a"}, false, []string{"a", "b", "c"}},
		{-1, []string{"b", "c", "a"}, true, []string{"c", "b", "a"}},
		{0, []string{"b", "c", "a"}, false, []string{"a", "b", "c"}},
		{0, []string{"b", "c", "a"}, true, []string{"c", "b", "a"}},
		{1, []string{"b", "c", "a"}, false, []string{"a", "b", "c"}},
		{1, []string{"b", "c", "a"}, true, []string{"c", "b", "a"}},
		{2, []string{"b", "c", "a"}, false, []string{"a", "b", "c"}},
		{2, []string{"b", "c", "a"}, true, []string{"c", "b", "a"}},
		{1, []string{"b b", "c a", "a c"}, false, []string{"c a", "b b", "a c"}},
		{1, []string{"b a", "c c", "a b"}, true, []string{"c c", "a b", "b a"}},
		{4, []string{"b b", "c a", "a c"}, false, []string{"a c", "b b", "c a"}},
		{4, []string{"b a", "c b ", "a c"}, true, []string{"c b ", "b a", "a c"}},
		{0, []string{"a a", "c b", "a c"}, false, []string{"a a", "a c", "c b"}},
		{0, []string{""}, true, []string{""}},
		{0, []string{"a", "b b", "c"}, false, []string{"a", "b b", "c"}},
		{1, []string{"a", "b b", "c"}, false, []string{"a", "c", "b b"}},
	}
	for _, table := range tables {
		result := sortByColumn(table.numberColumn, table.lines, table.neededToReverse)

		inputData := fmt.Sprintf("number of column = %v, lines = %v, need to reverse = %v", table.numberColumn, strings.Join(table.lines, "\n"), table.neededToReverse)
		assert.Equal(t, result, table.res, inputData)
	}
}
