package main

import (
	"fmt"
)

func main() {
	q := []string{"altoha", "+", "straponchik", "=", "ggay"}

	q[1] = q[len(q)-1]
	q = q[:len(q)-1]

	fmt.Println(q)
}
