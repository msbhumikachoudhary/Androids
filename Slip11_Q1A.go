// to check whether the accepted number is two digit or not.

package main
import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	if num >= 10 && num <= 99 || num <= -10 && num >= -99 {
		fmt.Println(num, "is a two-digit number.")
	} else {
		fmt.Println(num, "is not a two-digit number.")
	}
}