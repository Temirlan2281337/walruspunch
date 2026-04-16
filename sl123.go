package main

import "fmt"

func main() {
	var arr []int
	massiv := make([]int, 0)
	fmt.Println(len(massiv))
	// massiv = append(massiv, 1, 2, 3, 5)
	// fmt.Println(massiv)
	// fmt.Println(massiv[1 : len(massiv)-1])
	// if arr == nil {
	// 	fmt.Println("err")
	// } else {
	// 	fmt.Println("ne err")
	// }
	// if massiv == nil {
	// 	fmt.Println("err massiv")
	// } else {
	// 	fmt.Println("ne error mass")
	// }
	if arr == nil {
		fmt.Println("arr is NIL")
	}

	if massiv == nil {
		fmt.Println("massiv is NIL")
	} else {
		fmt.Println("massiv is NOT nil")
	}
	fmt.Println(arr)
}
