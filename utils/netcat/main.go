package main

import (
	"os"

	"github.com/Kirill-Mozes/mynetcat/funcs"
)

func main() {
	if len(os.Args) != 4 {
		println("more or less data")
	}

	host, port, protocol := os.Args[1], os.Args[2], os.Args[3]
	nc := funcs.NewNetcater(host, port, protocol)
	if err := nc.Start(); err != nil {
		println(err.Error())
	}
}
