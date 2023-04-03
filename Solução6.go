package main

import "fmt"

func main() {
	//criando o array de 10 elementos
	array := [10]int{7, 9, 10, 2, 5, 11, 6, 8, 3, 15}

	//inserindo um valor

	var valor int
	fmt.Println("Digite um valor para verificar a existencia dele no array:")
	fmt.Scanln(&valor)

	//verificando a existencia do elemento dentro do array
	encontrado := false
	for _, v := range array {
		if v == valor {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Printf("O valor %d está no array", valor)
	} else {
		fmt.Printf("O valor %d não está no array", valor)
	}
}
