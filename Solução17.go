package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum int

	for i := 0; i < len(array); i += 2 {
		sum += array[i]
	}

	fmt.Println("A soma dos elementos nas posições pares é: ", sum)

}
