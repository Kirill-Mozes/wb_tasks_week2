package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var group sync.WaitGroup
	c := make(chan interface{})
	for i := range channels {
		group.Add(1)
		go func(channel <-chan interface{}) {
			for val := range channel {
				c <- val
			}
			group.Done()
		}(channels[i])
	}
	go func() {
		group.Wait()
		close(c)
	}()

	return c
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
