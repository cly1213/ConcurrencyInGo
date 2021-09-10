package main

import "fmt"

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				swapNumbers(array, j)
			}
		}
	}
}

func swapNumbers(a []int, j int) {
	var temp int
	temp = a[j]
	a[j] = a[j+1]
	a[j+1] = temp
}

func main() {
	// The program should prompt the user to type in a sequence of up to 10 integers.
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	// The program should print the integers out on one line, in sorted order, from least to greatest.
	BubbleSort(input)
	fmt.Println("AFTER:  ", input)
}
