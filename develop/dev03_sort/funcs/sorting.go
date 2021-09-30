package funcs

import (
	"errors"
	"strconv"
	"strings"
)

type Sorting struct {
	text         []string
	flags        []string
	targetColumn int
	reverse      bool
	isSorted     bool
}

func InitSorting(lines []string, flags []string) (*Sorting, error) {
	if lines != nil {
		return &Sorting{lines, flags, -1, false, false}, nil
	}
	return nil, errors.New("empty line")
}

func (s *Sorting) GetRes() string {
	return strings.Join(s.text, "\n")
}

func (s *Sorting) Start() error {
	err := s.handlingFlags()
	if err != nil {
		return err
	}

	if len(s.flags) >= 2 {
		return errors.New("incompatible flags: " + strings.Join(s.flags, " "))
	} else if len(s.flags) == 0 {
		if s.targetColumn != -1 {
			sortedLines := sortByColumn(s.targetColumn, s.text, s.reverse)
			s.text = sortedLines
		} else {
			s.text = simpleSort(s.text, s.reverse)
		}
	} else {
		lastFlag := s.flags[0]
		switch lastFlag {
		case "-n":
			sortedText := sortByNumber(s.targetColumn, s.text, s.reverse)
			s.text = sortedText

		case "-h":
			sortedText := sortByNumberWithSuffix(s.targetColumn, s.text, s.reverse)
			s.text = sortedText

		case "-M":
			sortedText := sortByMonth(s.targetColumn, s.text, s.reverse)
			s.text = sortedText

		default:
			return errors.New("unknown flag")
		}
	}
	return nil
}

func isEqual(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, val := range lhs {
		if val != rhs[i] {
			return false
		}
	}
	return true
}

func (s *Sorting) handlingFlags() error {
	var flags []string
	for i := 0; i < len(s.flags); i++ {
		flag := s.flags[i]
		switch flag {
		case "-k":
			if i+1 < len(s.flags) {
				column, err := strconv.Atoi(s.flags[i+1])
				if err != nil {
					return err
				}
				s.targetColumn = column - 1
				i++
			} else {
				return errors.New("input number of column")
			}
		case "-r":
			s.reverse = true
		case "-u":
			uniqueLines := dropDuplicates(s.text)
			s.text = uniqueLines
		case "-b":
			for _, value := range s.text {
				value = strings.TrimLeft(value, " ")
			}
		case "-c":
			if isEqual(s.text, simpleSort(s.text, false)) {
				s.isSorted = true
			} else {
				s.isSorted = false
			}
		default:
			flags = append(flags, flag)
		}
	}
	s.flags = flags
	return nil
}
