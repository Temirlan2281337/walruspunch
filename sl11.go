package main

import (
	"fmt"
	"slices"
)

func main() {
	q := []string{"altoha", "+", "straponchik", "=", "ggay"}
	for i := 0; i < len(q); i++ {
		if q[i] == "+" {
			q = slices.Delete(q, i, i+1)
		}
	}
	fmt.Println(q)
}
