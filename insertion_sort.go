package main

import "fmt"

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			swap(j-1, j, array)
			j -= 1
		}
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 2, 5, 4, 1}
	fmt.Println(insertionSort(array))
}

/* output
[1 2 4 5 8]
*/
