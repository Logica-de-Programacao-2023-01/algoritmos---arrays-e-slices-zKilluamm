package main

import "fmt"

func main() {
	array := [10]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	var slice []float64

	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}
	fmt.Println(slice)

}
