package main

import "fmt"

func bubbleSort(arr []int) {
  len_array := len(arr)
  for i := 0; i < len_array -1 ; i++ {
    for j := 0; j < len_array - i -1 ; j++ {
      if arr[j] > arr[j+1] {
        arr[j] , arr[j+1] = arr[j+1], arr[j]
      }
    }
  }
  fmt.Println("after sort: " , arr)
}

func main() {
  bubbleSort([]int{5,9,4,2,1})
}
