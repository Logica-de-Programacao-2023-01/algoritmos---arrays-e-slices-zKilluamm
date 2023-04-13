package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var index1, index2 int
	fmt.Println("Digite o indice do primeiro número: ")
	fmt.Scan(&index1)
	fmt.Println("Digite o indice do segundo número número: ")
	fmt.Scan(&index2)

	//trocando de posições
	slice[index1], slice[index2] = slice[index2], slice[index1]

	fmt.Println("Slice resultante", slice)
}
