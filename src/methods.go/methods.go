package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int
	currency int
}


func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s%d ", e.name, e.currency, e.salary)
}


func main() {
	emp1 := Employee {
		name: "Sam Adolf",
		salary:  5000,
		currency: 2,
	}
	emp1.displaySalary()
}
