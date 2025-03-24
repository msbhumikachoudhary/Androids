//to swap two numbers using call by reference concept 

package main
import "fmt"

func swap(x, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func main() {
  var p, q int
  fmt.Print("Enter the two numbers to swap:")
  fmt.Scan(&p, &q)
  fmt.Println("Before swap:", "p =", p, "q =", q)
  swap(&p, &q)
  fmt.Println("After swap:", "p =", p, "q =", q)
}