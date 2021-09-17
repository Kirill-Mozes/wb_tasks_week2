package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseStringWithSuffixInNumber(t *testing.T) {
	tables := []struct {
		s   string
		res float64
	}{
		{"1k", 1024},
		{"1K", 1024},
		{"2K", 2048},
		{"10K", 10240},

		{"1kb", 1024},
		{"1KB", 1024},
		{"2KB", 2048},
		{"10KB", 10240},

		{"1m", 1024 * 1024},
		{"1M", 1024 * 1024},
		{"2M", 2 * 1024 * 1024},
		{"10M", 10 * 1024 * 1024},

		{"1mb", 1024 * 1024},
		{"1MB", 1024 * 1024},
		{"2MB", 2 * 1024 * 1024},
		{"10MB", 10 * 1024 * 1024},

		{"1g", 1024 * 1024 * 1024},
		{"1G", 1024 * 1024 * 1024},
		{"2G", 2 * 1024 * 1024 * 1024},
		{"10G", 10 * 1024 * 1024 * 1024},

		{"1gb", 1024 * 1024 * 1024},
		{"1GB", 1024 * 1024 * 1024},
		{"2GB", 2 * 1024 * 1024 * 1024},
		{"10GB", 10 * 1024 * 1024 * 1024},

		{"1t", 1024 * 1024 * 1024 * 1024},
		{"1T", 1024 * 1024 * 1024 * 1024},
		{"2T", 2 * 1024 * 1024 * 1024 * 1024},
		{"10T", 10 * 1024 * 1024 * 1024 * 1024},

		{"1tb", 1024 * 1024 * 1024 * 1024},
		{"1TB", 1024 * 1024 * 1024 * 1024},
		{"2TB", 2 * 1024 * 1024 * 1024 * 1024},
		{"10TB", 10 * 1024 * 1024 * 1024 * 1024},
	}
	for _, table := range tables {
		result, _ := parseStringWithSuffixInNumber(table.s)

		assert.Equal(t, result, table.res)
	}
}

func Test_sortByNumberWithSuffix(t *testing.T) {
	tables := []struct {
		numberColumn    int
		lines           []string
		neededToReverse bool
		res             []string
	}{
		{-1, []string{"1G", "1K", "2MB"}, false, []string{"1K", "2MB", "1G"}},
		{-1, []string{"1G", "1K", "2MB"}, true, []string{"1G", "2MB", "1K"}},
		{-1, []string{"2K", "1kb", "3KB"}, false, []string{"1kb", "2K", "3KB"}},
		{-1, []string{"2K", "1kb", "3KB"}, true, []string{"3KB", "2K", "1kb"}},
		{0, []string{"2K", "1kb", "3KB"}, false, []string{"1kb", "2K", "3KB"}},
		{0, []string{"2K", "1kb", "3KB"}, true, []string{"3KB", "2K", "1kb"}},
		{1, []string{"2K", "1kb", "3KB"}, false, []string{"1kb", "2K", "3KB"}},
		{1, []string{"2K", "1kb", "3KB"}, true, []string{"3KB", "2K", "1kb"}},
		{-1, []string{"2K", "1kb", "a"}, false, []string{"a", "1kb", "2K"}},
		{-1, []string{"2K", "1kb", "a"}, true, []string{"2K", "1kb", "a"}},
		{0, []string{"2K", "1kb", "a"}, false, []string{"a", "1kb", "2K"}},
		{0, []string{"2K", "1kb", "a"}, true, []string{"2K", "1kb", "a"}},
		{1, []string{"a 2K", "c 1kb", "b a"}, false, []string{"b a", "c 1kb", "a 2K"}},
		{1, []string{"a 2K", "c 1kb", "b a"}, true, []string{"a 2K", "c 1kb", "b a"}},
	}
	for _, table := range tables {
		result := sortByNumberWithSuffix(table.numberColumn, table.lines, table.neededToReverse)

		inputData := fmt.Sprintf("number of column = %v, lines = %v, need to reverse = %v", table.numberColumn, strings.Join(table.lines, "\n"), table.neededToReverse)
		assert.Equal(t, result, table.res, inputData)
	}
}
