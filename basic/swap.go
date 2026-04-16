package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	tempA := *a
	*a = *b
	*b = tempA
	// *a, *b = *b, *a (2 вариант решения без tempA)
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
