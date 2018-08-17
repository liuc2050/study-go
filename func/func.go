package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func main() {
	//var k [3]int = [...]int{1, 2, 3}
	var kkk []int = []int{1, 2, 3, 4}
	fmt.Printf("&kkk=%p\n", &kkk)
	testval(kkk)

	s := "12"
	fmt.Println(len(s[3:3]))
	EmployeeByID(12).Salary = 0
}

func testval(val []int) {
	fmt.Printf("&val=%p\n", &val)
}

func EmployeeByID(_ int) *Employee {
	var c Employee
	return &c
}
