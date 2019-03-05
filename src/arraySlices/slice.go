package main

import (
	"fmt"
)

func main() {
	a := [5]int{76,77,78,79,80}
	var b []int = a[1:4] // creates a slice from
	fmt.Println(b)
	fmt.Println(a)

	main2()
}

func main2() {
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
}
