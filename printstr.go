package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	fmt.Println("")
}

func main() {
	PrintStr("Hello World!")
}
