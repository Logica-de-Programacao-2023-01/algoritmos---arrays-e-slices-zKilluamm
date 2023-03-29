package main

import "fmt"

func main() {

	//criando as variaveis

	var numeros = [3]int{5, 10, 20}
	fmt.Println("Os valores armazenados no array são", numeros)

	//calculando a soma dos elementos dentro do array

	soma := 0

	for _, valor := range numeros {
		soma += valor
	}
	fmt.Printf("A soma dos valores que foram armazenados no array é %d", soma)

}
