package main

import (
	"fmt"
)

func main() {
	a := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "Пятка", "aaa"}
	res := SearchAnagramm(&a)
	fmt.Println(*res)

	b := []string{"", "", "", "п", "п", "ба", "аб"}
	res = SearchAnagramm(&b)
	fmt.Println(*res)
}
