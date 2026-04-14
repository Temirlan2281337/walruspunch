package main

import "fmt"

func FirstRune(s string) rune {
	res:=  rune(s[i])
for i:=0 ; i < len(s); i++{
	res = res + (rune(s[0]))
}
return res
}
func main() {
	fmt.Print(FirstRune("Hello!"))
	fmt.Print(FirstRune("Salut!"))
	fmt.Print(FirstRune("Ola!"))
	
}