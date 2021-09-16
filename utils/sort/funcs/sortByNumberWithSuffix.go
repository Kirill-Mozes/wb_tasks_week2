package funcs

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

var suffixes = map[string]int{
	"KB": 1 << 10,
	"K":  1 << 10,
	"MB": 1 << 20,
	"M":  1 << 20,
	"GB": 1 << 30,
	"G":  1 << 30,
	"TB": 1 << 40,
	"T":  1 << 40,
}

func parseStringWithSuffixInNumber(s string) (float64, error) {
	s = strings.ToUpper(s)
	for key, val := range suffixes {
		i := strings.Index(s, key)
		if i > -1 {
			numberString := s[:i]
			number, err := strconv.ParseFloat(numberString, 64)
			if err != nil {
				return 0, err
			}
			numberInBytes := number * float64(val)
			return numberInBytes, nil
		}
	}
	return 0, errors.New("wrong number")
}

func sortByNumberWithSuffix(numberColumn int, lines []string, neededToReverse bool) []string {
	var keys []float64
	keyColumnToLineMap := map[float64][]string{}
	var result []string
	var notNumber []string

	for _, line := range lines {
		sortedField, searched := getFieldForSortFromLines(numberColumn, line)
		if !searched {
			notNumber = append(notNumber, line)
		} else {
			key, err := parseStringWithSuffixInNumber(sortedField)
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

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		sort.Sort(sort.Reverse(sort.StringSlice(notNumber)))

		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		result = append(result, notNumber...)
	} else {
		sort.Strings(notNumber)
		sort.Float64s(keys)

		result = append(result, notNumber...)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}

	return result
}
