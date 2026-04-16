package main

import (
	"fmt"
)

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i++ {
		runes[i], runes[j] = runes[j], runes[i]
		j--
	}
	return string(runes)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
