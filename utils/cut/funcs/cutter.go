package funcs

import (
	"encoding/json"
	"errors"
	"strings"
)

type Cuter struct {
	text      []string
	flags     []string
	result    []string
	fields    []int
	delimiter string
	separated bool
}

func InitCut(lines []string, flags []string) (*Cuter, error) {
	if lines == nil {
		return nil, errors.New("empty file")
	}
	return &Cuter{
		text:      lines,
		flags:     flags,
		delimiter: "\t",
	}, nil
}

func (s *Cuter) GetRes() string {
	return strings.Join(s.result, "\n")
}

func (s *Cuter) Start() error {
	err := s.parseFlags()
	if err != nil {
		return err
	}
	if s.fields != nil {
		if s.delimiter == "" {
			s.result = s.text
		} else {
			s.result = myCut(s.text, s.delimiter, s.fields, s.separated)
		}
	} else {
		return errors.New("no fields")
	}

	return nil
}

func (g *Cuter) parseFlags() error {
	for i := 0; i < len(g.flags); i++ {
		flag := g.flags[i]
		switch flag {
		case "-f":
			if i+1 >= len(g.flags) {
				return errors.New("no argument for flag -f")
			}
			fieldsSliceString := "[" + g.flags[i+1] + "]"
			var fieldsIntSlice []int
			if err := json.Unmarshal([]byte(fieldsSliceString), &fieldsIntSlice); err != nil {
				panic(err)
			}
			i++
			g.fields = fieldsIntSlice

		case "-d":
			if i+1 >= len(g.flags) {
				return errors.New("no argument for flag -d")
			}
			g.delimiter = g.flags[i+1]
			i++
		case "-s":
			g.separated = true
		default:
			return errors.New("incorret input")
		}
	}
	return nil
}
