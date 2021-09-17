package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortByMonth(t *testing.T) {
	tables := []struct {
		numberColumn    int
		lines           []string
		neededToReverse bool
		res             []string
	}{
		{-1, []string{"jun", "may", "aug"}, false, []string{"may", "jun", "aug"}},
		{-1, []string{"jun", "may", "aug"}, true, []string{"aug", "jun", "may"}},
		{-1, []string{"jun a", "may a", "aug a"}, false, []string{"aug a", "jun a", "may a"}},
		{-1, []string{"jun a", "may a", "aug a"}, true, []string{"may a", "jun a", "aug a"}},
		{0, []string{"jun a", "may a", "aug a"}, false, []string{"may a", "jun a", "aug a"}},
		{0, []string{"jun a", "may a", "aug a"}, true, []string{"aug a", "jun a", "may a"}},
		{1, []string{"a jun", "a may", "a aug"}, false, []string{"a may", "a jun", "a aug"}},
		{1, []string{"a jun", "a may", "a aug"}, true, []string{"a aug", "a jun", "a may"}},
		{1, []string{"a jun", "a may", "b a", "a a"}, false, []string{"a a", "b a", "a may", "a jun"}},
		{1, []string{"a jun", "a may", "a b", "a a"}, true, []string{"a jun", "a may", "a b", "a a"}},
	}
	for _, table := range tables {
		result := sortByMonth(table.numberColumn, table.lines, table.neededToReverse)

		inputData := fmt.Sprintf("number of column = %v, lines = %v, need to reverse = %v", table.numberColumn, strings.Join(table.lines, "\n"), table.neededToReverse)
		assert.Equal(t, result, table.res, inputData)
	}
}
