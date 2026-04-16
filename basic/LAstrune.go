package main

import "fmt"

func main() {
	fmt.Printf("%c", LastRune("Hello!"))
	fmt.Printf("%c", LastRune("Salut!"))
	fmt.Printf("%c", LastRune("Ola!"))
	fmt.Printf("%c", '\n')
}

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]
}
