package main

import "fmt"

func main() {
	numbers := make([]int, 0, 3)

	// numbers[0] = 100
	// numbers[2] = 200
	// numbers[3] = 300

	numbers = append(numbers, 400)
	fmt.Println(numbers, len(numbers), cap(numbers))
}
