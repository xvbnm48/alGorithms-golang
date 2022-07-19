package main

import "fmt"

func insertionSort(arr []int) []int {
  for j := 1; j < len(arr); j++ {
    key := arr[j]
    i := j - 1
    for i > -1 && arr[i] > key {
      arr[1+i] = arr[i]
      i -= 1
    }
    arr[i+1] = key
  }

  return arr
}

func main() {
  angka := []int{7,8,6,5,1,2}
  fmt.Println("sebelum di sorting: " , angka)
  fmt.Println("sesudah di sorting: " , insertionSort(angka))
}
