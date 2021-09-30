package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simpleSort(t *testing.T) {
	tables := []struct {
		lines           []string
		neededToReverse bool
		res             []string
	}{
		{[]string{"b", "c", "a"}, false, []string{"a", "b", "c"}},
		{[]string{"b", "c", "a"}, true, []string{"c", "b", "a"}},
		{[]string{"ba", "ca", "aa"}, false, []string{"aa", "ba", "ca"}},
		{[]string{"ba", "ca", "aa"}, true, []string{"ca", "ba", "aa"}},
		{[]string{"b a", "c a", "a a"}, false, []string{"a a", "b a", "c a"}},
		{[]string{"b a", "c a", "a a"}, true, []string{"c a", "b a", "a a"}},
	}
	for _, table := range tables {
		result := simpleSort(table.lines, table.neededToReverse)

		inputData := fmt.Sprintf("lines = %v, need to reverse = %v", strings.Join(table.lines, "\n"), table.neededToReverse)
		assert.Equal(t, result, table.res, inputData)
	}
}
