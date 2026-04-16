package main

import "fmt"

func StrLen(s string) int {
	count := 0
	// Цикл работает, пока i меньше длины строки
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
