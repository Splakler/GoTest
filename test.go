package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)

	fmt.Println("numbers[1:4] ==", numbers[1:4])

	numbers2 := numbers[3:6:7]
	printSlice(numbers2)

	numbers = append(numbers, 30000)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}
