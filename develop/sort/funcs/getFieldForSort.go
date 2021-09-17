package funcs

import "strings"

func getFieldForSortFromLines(numberColumn int, line string) (string, bool) {
	if line == "" {
		return "", false
	}

	if numberColumn < 0 {
		return line, true
	} else {
		lineBySpace := strings.Fields(line)
		if numberColumn < len(lineBySpace) {
			return lineBySpace[numberColumn], true
		} else {
			return "", false
		}
	}
}
