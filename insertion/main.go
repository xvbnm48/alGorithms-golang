package main

import "fmt"

func insertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i > -1 && arr[i] > key {
			arr[i+1] = arr[i]
			i -= 1
		}
		arr[i+1] = key
	}
	return arr
}

func main() {
	angka := []int{5, 4, 3, 2, 1}
	fmt.Println("before sort: ", angka)
	fmt.Println("after sort: ", insertionSort(angka))
}
