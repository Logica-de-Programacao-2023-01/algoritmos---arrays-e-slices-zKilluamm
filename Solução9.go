package main

import "fmt"

func main() {
	array := [6]float64{1.0, 2.5, 3.2, 4.7, 5.1, 6.0}
	var num float64
	fmt.Print("Digite um número para ser adicionado ao array: ")
	fmt.Scanln(&num)

	for i := 0; i < len(array); i++ {
		array[i] += num
	}
	fmt.Println(array)
}
