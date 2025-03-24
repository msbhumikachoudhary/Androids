//to illustrate the concept of 
//returning multiple values from a function. ( Add, Subtract, 
//Multiply, Divide)

package main

import "fmt"

func calc(a, b int) (int, int, int, int) {
    sum := a + b
		diff := a - b
		multi := a * b
		div := a / b
    return sum, diff, multi, div
}

func main() {
    var n1, n2 int

    fmt.Print("Enter the two numbers to add, subtract, multiply and divide:")
    fmt.Scanln(&n1, &n2)
    sum, diff, multi, div := calc(n1, n2)
    
		fmt.Println("Sum:", sum)
		fmt.Println("Difference:", diff)
	  fmt.Println("Multiplication:", multi)
		fmt.Println("Division:", div)
}