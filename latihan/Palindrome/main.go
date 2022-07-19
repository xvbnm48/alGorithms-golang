package main

import (
  "fmt"
  "bufio"
  "os"
)

func Palindrome(input string) bool {
  for i := 0; i < len(input)/2; i++ {
    if input[i] != input[len(input)-i-1] {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println("masukkan  nama: ")
  b := bufio.NewScanner(os.Stdin)
  b.Scan()
  result := Palindrome(b.Text())

  if result == true {
    fmt.Println("Palindrome")
  } else {
    fmt.Println("not Palindrome")
  }
}
