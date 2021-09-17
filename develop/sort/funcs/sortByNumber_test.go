package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortByNumber(t *testing.T) {
	tables := []struct {
		numberColumn    int
		lines           []string
		neededToReverse bool
		res             []string
	}{
		{-1, []string{"3", "1", "2"}, false, []string{"1", "2", "3"}},
		{-1, []string{"3", "1", "2"}, true, []string{"3", "2", "1"}},
		{-1, []string{"3 a", "1 a", "2 a"}, false, []string{"1 a", "2 a", "3 a"}},
		{-1, []string{"3 a", "1 a", "2 a"}, true, []string{"3 a", "2 a", "1 a"}},
		{0, []string{"3 a", "1 b", "2 c"}, false, []string{"1 b", "2 c", "3 a"}},
		{0, []string{"3 a", "1 b", "2 c"}, true, []string{"3 a", "2 c", "1 b"}},
		{1, []string{"3 b", "1 a", "2 c"}, false, []string{"1 a", "2 c", "3 b"}},
		{1, []string{"3 b", "1 a", "2 c"}, true, []string{"3 b", "2 c", "1 a"}},
		{0, []string{"d", "c", "a 2", "b 1"}, false, []string{"a 2", "b 1", "c", "d"}},
		{0, []string{"d", "c", "a 2", "b 1"}, true, []string{"d", "c", "b 1", "a 2"}},
		{1, []string{"d", "c", "a 2", "b 1"}, false, []string{"c", "d", "b 1", "a 2"}},
		{1, []string{"d", "c", "a 2", "b 1"}, true, []string{"a 2", "b 1", "d", "c"}},
		{3, []string{"d", "c", "a 2", "b 1"}, false, []string{"a 2", "b 1", "c", "d"}},
		{3, []string{"d", "c", "a 2", "b 1"}, true, []string{"d", "c", "b 1", "a 2"}},
		{4, []string{"d", "c", "a 2", "b 1"}, false, []string{"a 2", "b 1", "c", "d"}},
		{4, []string{"d", "c", "a 2", "b 1"}, true, []string{"d", "c", "b 1", "a 2"}},
	}
	for _, table := range tables {
		result := sortByNumber(table.numberColumn, table.lines, table.neededToReverse)

		inputData := fmt.Sprintf("number of column = %v, lines = %v, need to reverse = %v", table.numberColumn, strings.Join(table.lines, "\n"), table.neededToReverse)
		assert.Equal(t, result, table.res, inputData)
	}
}
