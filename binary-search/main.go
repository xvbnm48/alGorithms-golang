package main

import "fmt"

func search(nums []int, target int) int {
	middle, left, right := 0, 0, len(nums)-1

	for left <= right {
		middle = left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 10))
}
