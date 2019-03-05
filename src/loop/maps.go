package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 900
	fmt.Println("personSalary map contents: ", personSalary)

	accessingItems()
	deleteItems()
	lengthMap()
	assingn()
}

func assingn(){
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
}

func lengthMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("length is", len(personSalary))
}

func deleteItems() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 900
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
}


func accessingItems(){
	personSalary := map[string]int{
		"steve":12000,
		"jamie":15000,
	}
	personSalary["mike"] = 900
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
}
