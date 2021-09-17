package funcs

import (
	"sort"
	"strings"
)

func sortByColumn(numberColumn int, lines []string, neededToReverse bool) []string {
	keys := []string{}
	var notColumn []string
	keyColumnToLineMap := map[string][]string{}

	for _, line := range lines {
		sortedField, searched := getFieldForSortFromLines(numberColumn, line)
		if !searched {
			notColumn = append(notColumn, line)
		} else {
			key := strings.ToLower(sortedField)
			if _, ok := keyColumnToLineMap[key]; !ok {
				keys = append(keys, key)
			}
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
		}
	}

	var result []string

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(notColumn)))
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))

		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		result = append(result, notColumn...)
	} else {
		sort.Strings(notColumn)
		sort.Strings(keys)
		result = append(result, notColumn...)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}

	return result
}
