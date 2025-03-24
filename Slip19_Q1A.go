//to  illustrate  the  function returning multiple values(add, subtract). 

package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    sub := a - b
    return sum, sub
}

func main() {
    var n1, n2 int

    fmt.Print("Enter the two numbers to sum and subtraction:")
    fmt.Scanln(&n1, &n2)
    sum, sub := calculate(n1, n2)
    
		fmt.Println("Sum:", sum)
	  fmt.Println("Subtrcation:", sub)
}