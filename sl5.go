package main

import "fmt"

func main() {
	q := 123
	w := &q
	fmt.Println(*w)
	*w = 321
	fmt.Println(*w)
}
