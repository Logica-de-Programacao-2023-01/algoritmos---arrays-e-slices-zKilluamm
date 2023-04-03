package main

import "fmt"

func main() {
	// criando um array bidimencional
	matriz := [3][2]int{}

	//Inserção de dados pelo usuário
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor para os elementos [%d][%d] :", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}
	//imprimindo a matriz
	fmt.Println("matriz resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matriz[i][j])

		}

		fmt.Println()

	}

	fmt.Println("Essas são as atribuições das matrizes")
}
