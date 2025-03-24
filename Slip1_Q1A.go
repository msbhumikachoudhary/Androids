//to accept user choice and print answers using arithmetic operators.

package main

import "fmt"

func main() {
	var choice, num1, num2 int

	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case 2:
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case 3:
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error! Division by zero.")
		}
	default:
		fmt.Println("Invalid choice! Please select a valid operation.")
	}
}
