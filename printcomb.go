package main

import "fmt"

func main() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				fmt.Print(a)
				fmt.Print(b)
				fmt.Print(c)
				if !(a == 7 && b == 8 && c == 9) {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Print("\n")
}
