package main

import "fmt"

func main() {
	var q bool
	q = true
	w := &q
	fmt.Println(*w)
	*w = false
	fmt.Println(*w)
}
