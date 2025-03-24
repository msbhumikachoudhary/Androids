//to demonstrate function return multiple values.

package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    multi := a * b
    return sum, multi
}

func main() {
    var n1, n2 int

    fmt.Print("Enter the two numbers to sum and multiply:")
    fmt.Scanln(&n1, &n2)
    sum, multi := calculate(n1, n2)
    
		fmt.Println("Sum:", sum)
	  fmt.Println("Multiplication:", multi)
}
