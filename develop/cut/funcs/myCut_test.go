package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myCut(t *testing.T) {
	tables := []struct {
		lines                        []string
		delimiter                    string
		fields                       []int
		outputLinesOnlyWithSeparator bool
		res                          []string
	}{
		{[]string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}, " ", []int{2, 3}, false, []string{"2 3", "b c", "б в", ", !"}},
		{[]string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}, " ", []int{2, 3}, true, []string{"2 3", "b c", "б в", ", !"}},
		{[]string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}, "", []int{2, 3}, false, []string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}},
		{[]string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}, "", []int{2, 3}, true, []string{"1 2 3 4", "a b c d", "а б в г", ". , ! ?"}},
		{[]string{"1 2 3 4", "a,b,c,d", "а б в г", ". , ! ?"}, " ", []int{2, 3}, false, []string{"2 3", "a,b,c,d", "б в", ", !"}},
		{[]string{"1 2 3 4", "a,b,c,d", "а б в г", ". , ! ?"}, " ", []int{2, 3}, true, []string{"2 3", "б в", ", !"}},
		{[]string{"1 2 3 4", "a,b,c,d", "а б в г", ". : ! ?"}, ",", []int{2, 3}, false, []string{"1 2 3 4", "b,c", "а б в г", ". : ! ?"}},
		{[]string{"1 2 3 4", "a,b,c,d", "а б в г", ". : ! ?"}, ",", []int{2, 3}, true, []string{"b,c"}},
	}
	for _, table := range tables {
		result := myCut(table.lines, table.delimiter, table.fields, table.outputLinesOnlyWithSeparator)
		inputData := fmt.Sprintf("lines = \n(%v)\ndelimiter = (%v)\nfields = %v\nonly with separator = %v",
			strings.Join(table.lines, "\n"), table.delimiter, table.fields, table.outputLinesOnlyWithSeparator)
		assert.Equal(t, result, table.res, inputData)
	}
}
