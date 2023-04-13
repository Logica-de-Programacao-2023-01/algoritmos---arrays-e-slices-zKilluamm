package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num int
	fmt.Print("Digite os números: ")
	fmt.Scan(&num)

	array[0] += num
	array[len(array)-1] += num

	fmt.Println("Array resultante:", array)
}
