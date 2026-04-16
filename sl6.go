package main

import "fmt"

func main() {
	qwe := make([]string, 0)
	qwe = append(qwe, "qq", "and", "mtt")
	zxc := ""
	for i := 0; i < len(qwe); i++ {
		zxc = zxc + qwe[i] + " "
	}
	fmt.Println(zxc)
}
