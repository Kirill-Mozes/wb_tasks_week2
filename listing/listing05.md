Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {		// сравнивается {type:*customError, val: nil} и nil
		println("error")// напишет это, т.к. тип хранящего внутри обхекта не nil, a *customError
		return
	}
	println("ok")
}
```

Ответ: nil error
```
Для правильной работы test должен возвращать объекты типа интерфейс error

```
