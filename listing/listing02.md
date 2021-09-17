Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {{//defer может читать и присваивать в x - т.к. x - именованный возврат
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x // здесь defer не может поменять возвращаемое значение, т.к. не именованый возврат
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ: 2 1

```


defer вызывается после return но до момента получения значения вызывающей стороной

При работе с defer - 3 правила
1) Аргументы вычисляются в момент выполнения строки ex: defer Add(a,b)
2) Defer'ы - выполняются по стеку
3) Defer может читать/писать в именованный возврат

```