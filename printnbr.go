package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		if n <= -10 {
			PrintNbr(-(n / 10))
		}
		lastdigit := -(n % 10)
		z01.PrintRune('0' + rune(lastdigit))
		return
	}
	if n > 0 {
		if n >= 10 {
			PrintNbr(n / 10)
		}
		digit := n % 10
		z01.PrintRune(('0' + rune(digit)))
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
