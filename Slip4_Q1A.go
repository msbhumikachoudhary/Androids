// to print a recursive sum of digits of a given number. 

package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number:")
	fmt.Scanln(&input)

	sum := sumDig(input)
	fmt.Println("Sum of digits:", sum)
}

func sumDig(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumDig(n/10)
}