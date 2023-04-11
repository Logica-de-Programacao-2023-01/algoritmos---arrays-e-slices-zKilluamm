package main

import "fmt"

func main() {
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var x, y int
	fmt.Println("Digite um índice de linha: ")
	fmt.Scan(&x)
	fmt.Println("Digite um índice de coluna: ")
	fmt.Scan(&y)

	fmt.Printf("Valor armazenado na posição [%d][%d]: %d", x, y, array[x][y])
}
