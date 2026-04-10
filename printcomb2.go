package main

import "fmt"

func main() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			fmt.Print(a)
			fmt.Print(b)
			if !(a == 9 && b == 9) {
				fmt.Print(", ")
			}
		}
	}
	fmt.Print("\n")
}
