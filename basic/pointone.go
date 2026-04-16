package main

import "fmt"

func altoshka(n *int) {
	*n = 3243
}

func main() {
	bc := 0
	altoshka(&bc)
	fmt.Println(bc)
}
