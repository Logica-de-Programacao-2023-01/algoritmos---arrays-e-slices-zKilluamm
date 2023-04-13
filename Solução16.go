package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int

	for _, val := range arr {
		if val%2 == 0 {
			slice = append(slice, val)
		}
	}
	fmt.Println("Os números pares do array são: ", slice)
}
