package main

import (
	"fmt"
)

func main() {
	b := 255
	var a * int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println("b",b)
	changeValue()
}

func changeValue() {
	b := 255
	a := &b
	fmt.Println("address of b is",a)
	fmt.Println("value of b is ", *a)
	*a++
	fmt.Println("new value of b is",b)
}
