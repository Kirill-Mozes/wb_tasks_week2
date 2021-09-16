package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/Kirill-Mozes/mygrep/funcs"
)

func readStrings(source io.Reader) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(source)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}
	return lines, nil
}

func main() {

	desiderStr := os.Args[1]
	flags := os.Args[2:]
	stat, err := os.Stdin.Stat()
	if err != nil {
		println(err.Error())
		return
	}
	var forRead io.Reader
	if stat.Mode()&os.ModeNamedPipe != 0 {
		forRead = os.Stdin
	} else {
		if len(flags) == 0 {
			return
		}
		fileName := flags[0]
		flags = flags[1:]
		file, err := os.Open(fileName)
		//defer file.Close()
		if err != nil {
			println(err.Error())
			return
		}
		forRead = file
	}
	lines, err := readStrings(forRead)
	if err != nil {
		println(err.Error())
		return
	}
	greper, err := funcs.InitGrep(lines, flags, desiderStr)

	if err != nil {
		println(err.Error())
		return
	}
	err = greper.Start()
	if err != nil {
		println(err.Error())
		return
	}
	println(greper.GetRes())
}
