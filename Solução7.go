package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	var num int

	fmt.Print("DIgite um número inteiro:")
	fmt.Scanln(&num)

	for _, valor := range slice {
		if valor == num {
			fmt.Println("O número já está presente no slice")
			fmt.Println(slice)
			return
		}
	}

	slice = append(slice, num)
	fmt.Println(slice)
}
