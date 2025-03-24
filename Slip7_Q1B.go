// to create structure student. Write a method show() whose receiver is a pointer of struct student.

package main

import (
	"fmt"
)

type Student struct {
	RollNo int
	Name   string
	Marks  float64
}

func (s *Student) show() {
	fmt.Println("Student Details:")
	fmt.Println("Roll No:", s.RollNo)
	fmt.Println("Name:", s.Name)
	fmt.Println("Marks:", s.Marks)
}

func main() {
	student1 := Student{RollNo: 101, Name: "Alice", Marks: 89.5}
	student1.show()
}
