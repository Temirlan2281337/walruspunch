package main

import "fmt"

func main() {
	qwe := "supstring"
	rty := []rune(qwe)
	fmt.Println(string(rty[2:5]))
}
