package main

import "fmt"

func bubbleSort(arr []int) {
  iniArray := len(arr)

  for i:= 0; i < iniArray -1; i++ {
    for j := 0; j < iniArray -1 -i; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1] , arr[j]
      }
    }
  }
  fmt.Println("sesudah di sorting: " , arr)
}

func main() {
  contoh := []int{9,6,5,48,1}
  fmt.Println("sebelum di sort: ", contoh)
  bubbleSort(contoh)
}
