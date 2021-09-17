package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Наполнение канала значениями с случайными таймаутами
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// Слияние каналов
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			// Проблема в том, что когда каналы закроются,
			// они будут возвращать дефолтные значения,
			// для проверки, закрыт ли канал - можно посмотреть на второе значение
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func newMerge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			if a == nil && b == nil {
				close(c)
				return
			}
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := newMerge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/*
1
2
3
4
5
6
7
8
0
0
0
0
0
0
0
0  на бесконечность
*/
