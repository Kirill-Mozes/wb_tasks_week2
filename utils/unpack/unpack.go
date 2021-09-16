package main

import (
	"errors"
	"fmt"
)

//var ErrInvalidString = errors.New("invalid string")

func Unpacking(str string) (string, error) {

	d := make([]rune, 0, len(str))
	for _, k := range str {
		d = append(d, k)
	}
	res := make([]rune, 0, len(str))
	for i := 0; i < len(d); i++ { //идем по строке
		if (d[i] >= 'a' && d[i] <= 'z') || (d[i] >= 'A' && d[i] <= 'Z') { // если буква тогда все хорошо
			res = append(res, d[i])
			number := 0
			j := i + 1
			counter := 0 //количество цифр подряд
			for j < len(d) && d[j] >= '0' && d[j] <= '9' {
				if counter == 0 && int(d[j])-48 == 0 {
					return "", errors.New("invalid string")
				}
				number = number*10 + (int(d[j]) - 48)
				j++
				counter++
			}
			for k := 1; k < number; k++ {
				res = append(res, d[i])
			}
			i += counter
		} else if d[i] == '\\' && i+1 < len(d) {
			// пишем в res следующий эелемент
			res = append(res, d[i+1])
			number := 0
			j := i + 2
			counter := 0 //количество цифр подряд
			for j < len(d) && d[j] >= '0' && d[j] <= '9' {
				if counter == 0 && int(d[j])-48 == 0 {
					return "", errors.New("invalid string")
				}
				number = number*10 + (int(d[j]) - 48)
				j++
				counter++
			}
			for k := 1; k < number; k++ {
				res = append(res, d[i+1])
			}
			i++
			i += counter
		} else {
			fmt.Println(i)
			return "", errors.New("invalid string")
		}
	}
	return string(res), nil
}
