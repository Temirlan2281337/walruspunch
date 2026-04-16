package main

import "fmt"

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign *= (-1)
		start = 1
	} else if s[0] == '+' {
		sign = 1
		start = 1
	}

	for i := start; i < len(s); i++ {

		if !(s[i] >= '0' && s[i] <= '9') {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}
	return res * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi(""))
}
