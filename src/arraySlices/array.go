package main

import (
	"fmt"
)

func main() {
	var a  [3]int
	fmt.Println(a)
	arr()
}


func arr(){
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a{
		fmt.Println("%d the element of a is %.2f|n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
}
