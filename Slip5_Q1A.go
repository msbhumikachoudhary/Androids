

package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := "example.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, this is a text file!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created successfully:", fileName)
}
