package funcs

import (
	"sort"
	"strconv"
)

func sortByNumber(numberColumn int, lines []string, neededToReverse bool) []string {
	if numberColumn == -1 {
		numberColumn = 0
	}
	var keys []int
	var notNumber []string

	keyColumnToLineMap := map[int][]string{}
	for _, line := range lines {
		sortedField, searched := getFieldForSortFromLines(numberColumn, line)
		if !searched {
			notNumber = append(notNumber, line)
		} else {
			key, err := strconv.Atoi(sortedField)
			if err != nil {
				notNumber = append(notNumber, line)
			} else {
				if _, ok := keyColumnToLineMap[key]; !ok {
					keys = append(keys, key)
				}
				keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			}
		}
	}

	var result []string
	if neededToReverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		sort.Sort(sort.Reverse(sort.StringSlice(notNumber)))

		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		result = append(result, notNumber...)
	} else {
		sort.Ints(keys)
		sort.Strings(notNumber)

		result = append(result, notNumber...)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}

	return result
}
