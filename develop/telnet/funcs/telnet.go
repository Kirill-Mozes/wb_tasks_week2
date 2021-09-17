package funcs

import (
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Telneter struct {
	host    string
	port    string
	timeout time.Duration
}

func newTelnet(host, port string, timeout time.Duration) *Telneter {
	return &Telneter{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func Start(timeout *time.Duration, host, port string) error {
	t := newTelnet(host, port, *timeout)
	client := NewClient(t.host+":"+t.port, t.timeout, os.Stdin, os.Stdout)
	if err := client.BuildConnection(); err != nil {
		return err
	}
	defer client.Close()

	signalCh := make(chan os.Signal, 1)
	errorCh := make(chan error, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go Receive(client, errorCh)
	go Send(client, errorCh)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-signalCh:
				println("program trying to exit....")
				return
			case err := <-errorCh:
				if err != nil {
					if err == io.EOF {
						println("recieved EOF")
					} else {
						println("...Connection was closed by peer")
					}
					return
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()
	return nil
}

func Send(c *Client, errorCh chan error) {
	for {
		if err := c.Send(); err != nil {
			errorCh <- err
			return
		}
	}
}

func Receive(c *Client, errorCh chan error) {
	for {
		if err := c.Receive(); err != nil {
			errorCh <- err
			return
		}
	}
}
