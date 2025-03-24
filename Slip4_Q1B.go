// to sort array elements in ascending order.

package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{5, 3, 8, 1, 2, 7, 4, 6}

	fmt.Println("Original array:", arr)

	sort.Ints(arr)
	
  fmt.Println("Sorted array in ascending order:", arr)
}
