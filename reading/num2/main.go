package main

import (
	"fmt"
)

func test() (x int) {
	defer func() { //дефер выполняется после return и до получения данных 2й стороной
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
		//fmt.Println(x) // он выполнится но в зачет не идет
	}()
	x = 1
	return x // вернули х раньше чем выполняется дефер
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
