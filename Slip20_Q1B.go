//how to create a channel and 
//illustrate how to close a channel 
//using for range loop and close function.

package main
import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        for i := 1; i <= 10; i++ {
            ch <- i
        }
        close(ch)
    }()

    for num := range ch {
        fmt.Println(num)
    }
    fmt.Println("Channel Closed Successfully")
}