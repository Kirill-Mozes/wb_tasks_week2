package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"} //len:3, cap:3
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"         // просходит над нужным слайсом
	i = append(i, "4") //слайс переместился в другую обл памяти и там изменяется
	i[1] = "5"         //уже над другим
	i = append(i, "6")
}

//[3 2 3]
