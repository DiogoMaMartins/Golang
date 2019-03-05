package main

//import "fmt"
import (
	"fmt"
	"math"
)

func main() {
	var age int // is a variable declaration
	fmt.Println("my age is ", age)
	age = 24 // assignment
	fmt.Println("my age now is ", age)

	var ageI int = 29 // variable declaration with initial value
	var nameI string = "Diogo"

	fmt.Println("my name is", nameI ,"and i have ", ageI)

	var width, height int = 100, 50 // declaring multiple variable

	fmt.Println("width is ", width ,"height is ", height)

	var monster, redbull float64

	fmt.Println("monster coust is", monster , "redbull coust is" , redbull)
	monster = 1.75
	redbull = 2.50
	fmt.Println("new price for monster", monster , "new price for redbull" ,redbull)

	a,b := 20,30 // declare variables a and b

	fmt.Println("a is ", a, "b is", b)
	
	maths()
}

func maths() {
	a,b := 145.8, 543.8
	c := math.Min(a,b)
	fmt.Println("minimum value is", c)
}
