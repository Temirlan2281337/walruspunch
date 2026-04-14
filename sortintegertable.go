package main

import "fmt"

func SortIntegerTable(table []int) {
for i:=0; i<len(table); i++{
	for j:= i+1; j<len(table); j++{
		if table[i] > table[j]{
		table[i], table[j] = table[j], table[i]
		}

	}

}
}


func main() {
    nums := []int{10, -2, 3, 0, 5}
    SortIntegerTable(nums)
    fmt.Println(nums) // Ожидаем: [-2 0 3 5 10]
}