// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

// Problem: Given a list of n integers arr[0..(n-1)], determine the number of different pairs of elements within it which sum to k.
// If an integer appears in the list multiple times, each copy is considered to be different; that is, two pairs are considered different
// if one pair includes at least one array index which the other doesn't, even if they include the same values.
//
// Solution: use hash map to calculate the counts of each number
// Iterate over the map contents and see if k - each key is in the map
// if adds up to itself, use combination formula, otherwise just multiply the counts
// Since we double count by iterating over the same pairs twice, divide by 2 in the end
// Time complexity: O(n)
// Space complexity: O(n)
//
// Alternative solution:
// sort the array, use 2 pointers from start and end, compare to target and decrement or increment
// pointers accordingly. Treat the same values as one index but assign the weight equal to the number of occurrences.
// Time complexity: O(nlogn)
// Space complexity: O(1)
func numberOfWays(arr []int, k int) int {
  countMap := make(map[int]int, 0)
  
  for _, n := range arr {
    if val, exists := countMap[n]; exists {
      countMap[n] = val + 1
    } else {
      countMap[n] = 1
    }
  }
  
  sum := 0
  for n, nCount := range countMap {
    diff := k - n
    if diffCount, exists := countMap[diff]; exists {
      if diff == n {
        sum += diffCount * (diffCount - 1)
      } else {
        sum += nCount * diffCount
      }
    }
  }
  
	return sum / 2;
}

func main() {
  fmt.Println(numberOfWays([]int{1, 2, 3, 4, 3}, 6))
  fmt.Println(numberOfWays([]int{1, 5, 3, 3, 3}, 6))
  fmt.Println(numberOfWays([]int{1, 1, 1, 1, 1}, 2))
}
