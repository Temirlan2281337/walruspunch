package main

import "fmt"

func main() {
	q := make([]string, 0)
	q = append(q, "enrico", "cubbum")
	w := make([]string, len(q))
	copy(w, q)

	fmt.Println(q)
	fmt.Println(w)
}
