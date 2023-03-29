package main

import "fmt"

func main() {
	//criando um array de floats
	num1 := [4]float64{2.0, 3.0, 6.0, 4.0}

	//variavel produto

	produto := 1.0

	//iterando sobre os elementos do array

	for _, valor := range num1 {
		produto *= valor
	}

	//imprimindo o resultado do produto dos elementos do array

	fmt.Printf("O resultado do produto dos elementos é :%.2f", produto)

}
