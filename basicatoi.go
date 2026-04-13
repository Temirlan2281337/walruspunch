package main

import "fmt"

func BasicAtoi(s string) int {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
	}
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
