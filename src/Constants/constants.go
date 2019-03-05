package main

import (
	"fmt"
	"math"
)

func main() {
	//const a = 55 //allowed
	//a = 89 reassignment not allowed

	fmt.Println("Hello, playground")
	var a = math.Sqrt(4) //allowed
	//const b = math.Sqrt(4) // not allowed
	fmt.Println(a)

	const hello = "Hello World"

	fmt.Println("type %T value %V", hello, hello)

	const typedhello string = "Hello World"

	var defaultName = "Diogo"
	type myString string
	var customName myString = "Diogo"

}


