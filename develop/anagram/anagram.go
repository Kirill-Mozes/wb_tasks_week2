package main

import (
	//"fmt"
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
} //все ради сортировки слова

func Dedup(words []string) []string {
	unique := []string{}

	for _, word := range words {
		// If we alredy have this word, skip.
		if contains(unique, word) {
			continue
		}

		unique = append(unique, word)
	}

	return unique
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
} // удаление дубликатов

func SearchAnagramm(arr *[]string) *map[string][]string {
	buf := make(map[string][]string)
	for _, v := range *arr {
		if len(v) <= 2 {
			continue //пустые строки и буквы не анаграммы и по условию слова только русские
		}
		v := strings.ToLower(v)
		key := SortStringByCharacter(v)
		buf[key] = append(buf[key], v) // заполнили мапу анаграмм
	}
	outMap := make(map[string][]string)
	for _, value := range buf {
		if len(value) == 1 {
			continue // пропускаем единичные множества
		}
		outMap[value[0]] = append(outMap[value[0]], value[0])
		for i := 1; i < len(value); i++ {
			outMap[value[0]] = append(outMap[value[0]], value[i]) //заполнили выходную мапу
		}
	}
	for key, value := range outMap {
		sort.Strings(value)        // сортируем
		outMap[key] = Dedup(value) // дропаем дубликаты

	}
	return &outMap
}
