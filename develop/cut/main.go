package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/Kirill-Mozes/mycut/funcs"
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

	flags := os.Args[1:]
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
	cuter, err := funcs.InitCut(lines, flags)

	if err != nil {
		println(err.Error())
		return
	}
	err = cuter.Start()
	if err != nil {
		println(err.Error())
		return
	}
	println(cuter.GetRes())
}
