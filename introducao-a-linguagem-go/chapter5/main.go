package main

import "fmt"

func main() {
	numbers := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	var min int

	for i := 0; i < len(numbers); i++ {
		if i == 0 || numbers[i] < min {
			min = numbers[i]
		}
	}

	fmt.Println(min)
}
