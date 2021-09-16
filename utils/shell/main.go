package main

import (
	"os"

	"github.com/Kirill-Mozes/myshell/funcs"
)

func main() {
	commander, err := funcs.InitSheller(os.Stdout, os.Stdin)
	if err != nil {
		println(err.Error())
	}

	if err := commander.Start(); err != nil {
		println(err.Error())
	}
}
