// to print a multiplication table of number using function. 

package main
import "fmt"

func MultiTable(num int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}

func main() {
    var num int
    fmt.Print("Enter the number: ")
    fmt.Scanln(&num)
    MultiTable(num)
}