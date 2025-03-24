// to create a buffered channel, 
//store few values in it and find channel capacity and length. Read 
//values from channel and find modified length of a channel 

package main
import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 8
	ch <- 5
	ch <- 11
	ch <- 17
	ch <- 23
	fmt.Println("Capacity:", cap(ch)) 
	fmt.Println("Length:", len(ch))   
	val := <-ch
	fmt.Println("Modified length:", len(ch)) 
	fmt.Println("Read value:", val) 
}