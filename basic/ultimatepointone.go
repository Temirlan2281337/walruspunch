package main

import "fmt"

func hahahihi(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	hahahihi(&n)
	fmt.Println(a)
}
