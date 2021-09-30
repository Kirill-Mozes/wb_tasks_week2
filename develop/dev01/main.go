package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {

	options := ntp.QueryOptions{Timeout: 10 * time.Second, TTL: 15}
	response, err := ntp.QueryWithOptions("0.beevik-ntp.pool.ntp.org", options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = response.Validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	time := time.Now().Add(response.ClockOffset).Format("15:04:05")
	fmt.Fprintln(os.Stderr, time)
	os.Exit(0)
}
