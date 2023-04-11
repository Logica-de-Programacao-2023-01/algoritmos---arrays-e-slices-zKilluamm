package main

import "fmt"

func main() {
	array := [5]int{1, 5, 6, 4, 3}

	var mult3 []int

	for _, num := range array {
		if num%3 == 0 {
			mult3 = append(mult3, num)
		}
	}
	fmt.Println(mult3)

}
