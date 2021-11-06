// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

// Problem: Given two arrays A and B of length N, determine if there is a way to make A equal to B
// by reversing any subarrays from array B any number of times.
func areTheyEqual(arr1 []int, arr2 []int) bool {
  left := 0
  right := len(arr1) - 1
  
  for arr1[left] == arr2[left] {
    if left == right  {
      return true
    }
    left++
  }
  
  for arr1[right] == arr2[right] {
    if left == right {
      return true
    }
    right--
  }
  
  for left != right {
    if arr1[left] != arr2[right] {
      return false
    }
    left++
    right--
  }
  
	return true;
}

func main() {
  fmt.Println(areTheyEqual([]int{1, 2, 3, 4}, []int{1, 4, 3, 2}))
  fmt.Println(areTheyEqual([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
}
