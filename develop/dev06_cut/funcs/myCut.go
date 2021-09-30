package funcs

import (
	"sort"
	"strings"
)

func myCut(text []string, delimiter string, fields []int, separated bool) []string {
	if delimiter == "" {
		return text
	}
	var result []string
	sort.Ints(fields)
	for _, line := range text {
		if strings.Contains(line, delimiter) {
			lineSliceByDelimiter := strings.Split(line, delimiter)
			var newLine string
			for _, val := range fields {
				val--
				if val < len(lineSliceByDelimiter) {
					if newLine != "" {
						newLine += delimiter + lineSliceByDelimiter[val]
					} else {
						newLine += lineSliceByDelimiter[val]
					}
				}
			}
			result = append(result, newLine)
		} else if !separated {
			result = append(result, line)
		}
	}
	return result
}
