package main

import "fmt"

type customError struct {
	msg string
}

// реализация интерфейса error
func (e *customError) Error() string {
	return e.msg
}

// функция, которая возвращает ошибку
func test() *customError {
	//func test() error {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	//fmt.Printf("%T\n", err) // *main.customError
	fmt.Println(err) // nil
	if err != nil {  // сравнивается {type:*customError, val: nil} и nil
		println("error")
		return
	}
	println("ok")
}

// <nil>
// error
// как и в 3 присваивание и интерфейсы
