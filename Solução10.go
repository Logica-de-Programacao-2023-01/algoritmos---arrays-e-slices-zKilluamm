package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Ints(slice)
	fmt.Printf("Valor mínimo: %d\nValor máximo: %d\n", slice[0], slice[len(slice)-1])
}
