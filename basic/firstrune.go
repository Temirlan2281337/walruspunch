package main

import "fmt"

func FirstRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}

func main() {
	fmt.Printf("%c", FirstRune("Hello!"))
	fmt.Printf("%c", FirstRune("Salut!"))
	fmt.Printf("%c", FirstRune("Ola!"))
}
