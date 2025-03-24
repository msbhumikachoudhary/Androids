//using function to check whether accepts number is palindrome or not. 

package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	ori := num
	rev := 0

	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num = num / 10
	}

	return ori == rev
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome.")
	} else {
		fmt.Println(num, "is not a palindrome.")
	}
}
