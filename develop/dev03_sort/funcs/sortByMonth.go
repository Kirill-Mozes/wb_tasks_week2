package funcs

import (
	"sort"
	"strings"
)

type Month []string

var month = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func (m Month) Len() int           { return len(m) }
func (m Month) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Month) Less(i, j int) bool { return month[m[i]] < month[m[j]] }

func sortByMonth(numberColumn int, lines []string, neededToReverse bool) []string {
	var keys Month
	var notMonth []string
	keyColumnToLineMap := map[string][]string{}

	for _, line := range lines {
		sortedField, searched := getFieldForSortFromLines(numberColumn, line)
		if !searched {
			notMonth = append(notMonth, line)
		} else {
			key := strings.ToUpper(sortedField)
			if _, ok := month[key]; !ok {
				notMonth = append(notMonth, line)
			} else {
				if _, ok = keyColumnToLineMap[key]; !ok {
					keys = append(keys, key)
				}

				keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			}
		}
	}

	var result []string

	if neededToReverse {
		sort.Sort(sort.Reverse(keys))
		sort.Sort(sort.Reverse(sort.StringSlice(notMonth)))

		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		result = append(result, notMonth...)
	} else {
		sort.Strings(notMonth)
		sort.Sort(keys)

		result = append(result, notMonth...)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}

	return result
}
