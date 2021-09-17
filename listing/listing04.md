Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// после выполнения цикла канал останется открытым
		// close(ch) // для исправления
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
```
Программа напечатает числа от 0 до 9. Затем главная горутина зависнет в ожидании закрытия канала ch.
т.к. конструкция for range channel будет ожидать и читать значения, пока канал открыт.
```