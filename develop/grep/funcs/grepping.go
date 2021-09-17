package funcs

import (
	"errors"
	"strconv"
	"strings"
)

type Grepping struct {
	text            []string
	flags           []string
	desiderStrtring string
	searchedStringd []string
	linesAfter      int
	linesBefore     int
	needCount       bool
	count           int
	ignoreCase      bool
	invert          bool
	fixed           bool
	lineNum         bool
}

func InitGrep(lines []string, flags []string, desiderStr string) (*Grepping, error) {
	if lines == nil {
		return nil, errors.New("empty file")
	}
	return &Grepping{
		text:            lines,
		flags:           flags,
		desiderStrtring: desiderStr,
		fixed:           true,
	}, nil
}

func (s *Grepping) GetRes() string {
	if s.needCount {
		return strconv.Itoa(s.count)
	}
	return strings.Join(s.searchedStringd, "\n")
}

func (g *Grepping) parseFlags() error {
	for i := 0; i < len(g.flags); i++ {
		flag := g.flags[i]
		switch flag {
		case "-A":
			if i+1 >= len(g.flags) {
				return errors.New("no argument for flag -A")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			g.linesAfter = count
		case "-B":
			if i+1 >= len(g.flags) {
				return errors.New("no argument for flag -B")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			g.linesBefore = count
		case "-C":
			if i+1 >= len(g.flags) {
				return errors.New("no argument for flag -C")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			g.linesBefore, g.linesAfter = count, count
		case "-c":
			g.needCount = true
		case "-i":
			g.ignoreCase = true
		case "-v":
			g.invert = true
		case "-F":
			g.fixed = false
		case "-n":
			g.lineNum = true
		default:
			return errors.New("incorret input")
		}
	}
	return nil
}

func (g *Grepping) Start() error {
	err := g.parseFlags()
	if err != nil {
		return err
	}
	g.searchedStringd, g.count, err = manGrep(g.text, g.desiderStrtring, g.linesAfter, g.linesBefore, g.ignoreCase, g.lineNum, g.fixed, g.invert)
	if err != nil {
		return err
	}
	return nil
}
