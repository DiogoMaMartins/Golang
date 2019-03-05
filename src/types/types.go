package main

import (
	"fmt"
	"unsafe"
)

func main () {
	a := true
	b := false

	fmt.Println("a:", a, "b:" , b)

	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	size()
	floating()
	complexNumbers()
	strings()
	convertedToInt()
	assigned()
}


func assigned() {
	i := 10
	var j  float64 = float64(i)
	fmt.Println("j is ", j)
}

func convertedToInt() {
	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println("j is converted to int ", sum)
}

func size () {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Println("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))
}


func floating() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n",a,b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum , "diff", diff)

	no1, no2 := 56,89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}


func complexNumbers(){
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 +c2

	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

}

func strings() {
	first := "Diogo"
	last := "Martins"

	name := first + " "+ last
	fmt.Println("My name is ",name)
}

