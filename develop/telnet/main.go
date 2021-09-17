package main

import (
	"flag"
	"time"

	"github.com/Kirill-Mozes/mytelnet/funcs"
)

func main() {

	timeoutVar := flag.Duration("timeout", 10*time.Second, "connection timeout for client in seconds")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		println("usage: mytelnet --timeout=10s host port")
	}
	host := args[0]
	port := args[1]
	err := funcs.Start(timeoutVar, host, port)
	if err != nil {
		println(err.Error())
	}
}
