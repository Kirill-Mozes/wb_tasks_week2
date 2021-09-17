package funcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_manGrep(t *testing.T) {
	tables := []struct {
		lines        []string
		targerString string
		before       int
		after        int
		ignoreCase   bool
		lineNumber   bool
		regular      bool
		invert       bool
		resLines     []string
		resCount     int
	}{
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, false, false, []string{"cooler", "ert", "rty", "QWE"}, 2},

		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, false, false, []string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, false, false, []string{"1. cooler", "2. ert", "3. rty", "4. QWE"}, 2},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, true, false, []string{"cooler", "ert", "rty", "QWE"}, 2},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, false, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 4},

		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, false, false, []string{"1. cooler", "2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, true, false, []string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, false, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, true, false, []string{"1. cooler", "2. ert", "3. rty", "4. QWE"}, 2},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, false, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 4},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, true, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 4},

		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, true, false, []string{"1. cooler", "2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, false, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, true, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, true, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 4},

		{[]string{"cooler", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, true, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
	}
	for _, table := range tables {
		resultLines, resultCount, _ := manGrep(table.lines, table.targerString, table.after, table.before, table.ignoreCase, table.lineNumber, table.regular, table.invert)

		inputData := fmt.Sprintf("lines = (%v)\ntarget = (%v)\ncount before = %v, count after = %v\nignore case = %v, print line = %v, use regular = %v, invert = %v",
			strings.Join(table.lines, "\n"), table.targerString, table.before, table.after, table.ignoreCase, table.lineNumber, table.regular, table.invert)
		assert.Equal(t, resultLines, table.resLines, inputData)
		assert.Equal(t, resultCount, table.resCount, inputData)
	}
}
