package funcs

import (
	"regexp"
	"strconv"
	"strings"
)

func manGrep(textByLines []string, targerString string, after int, before int, ignoreCase bool, lineNumber bool, regular bool, invert bool) ([]string, int, error) {
	var result []string
	var count int
	lineAlreadyResult := map[int]struct{}{}

	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}

	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if regular {
			//println("pattern")
			matched, err := regexp.MatchString(targerString, line)
			if err != nil {
				return nil, 0, err
			}
			if invert {
				matched = !matched
				//println("invet")
			}
			if matched {
				//println("match")
				count++
				lenBefore := min(index, before)
				lenAfter := min(len(textByLines)-index-1, after)
				for i := index - lenBefore; i <= index+lenAfter; i++ {
					if lineNumber {
						textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
					}

					if _, contain := lineAlreadyResult[i]; !contain {
						result = append(result, textByLines[i])
						lineAlreadyResult[i] = struct{}{}
					}
				}
				//fmt.Println(count, " ", result)
			}
		} else {
			//println("Nopattern")
			contain := strings.Contains(line, targerString)
			if invert {
				contain = !contain
			}
			if contain {
				count++
				lenBefore := min(index, before)
				lenAfter := min(len(textByLines)-index-1, after)
				for i := index - lenBefore; i <= index+lenAfter; i++ {
					if lineNumber {
						textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
					}
					if _, alreadyIn := lineAlreadyResult[i]; !alreadyIn {
						result = append(result, textByLines[i])
						lineAlreadyResult[i] = struct{}{}
					}
				}
			}
		}
	}
	return result, count, nil
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
