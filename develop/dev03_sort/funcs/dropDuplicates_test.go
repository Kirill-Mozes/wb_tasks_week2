package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dropDuplicates(t *testing.T) {
	tables := []struct {
		lines []string
		res   []string
	}{
		{[]string{"abcd", "abcd"}, []string{"abcd"}},
		{[]string{"abcd"}, []string{"abcd"}},
		{[]string{"", ""}, []string{""}},
		{[]string{"abcd", "jlka"}, []string{"abcd", "jlka"}},
		{[]string{"abcd", "abcd", "jlka"}, []string{"abcd", "jlka"}},
		{nil, []string{}},
	}
	for _, table := range tables {
		result := dropDuplicates(table.lines)

		assert.Equal(t, result, table.res)
	}
}
