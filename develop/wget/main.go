package main

import (
	"os"

	"github.com/Kirill-Mozes/mywget/funcs"
)

func main() {
	if len(os.Args) != 2 {
		println("incorrect input")
		return
	}
	flags := os.Args[1:]
	url := flags[0]
	fileName := ""
	if len(flags) > 1 {
		fileName = flags[1]
	}
	wgeter, err := funcs.InitWget(url, fileName)

	if err != nil {
		println(err.Error())
		return
	}
	err = wgeter.Start()
	if err != nil {
		println(err.Error())
		return
	}

}
