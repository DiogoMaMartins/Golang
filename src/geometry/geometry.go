//geometry.go
package main

import (
	"fmt"
	"geometry/rectangle"
)


func main() {
	fmt.Println("Geometrical shape properties")

	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used */

	fmt.Printf("area of reactangle %.2f\n", reactangle.Area(rectLen,recWidth))

	/*Diagonal function of rectangle package used */

	fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}



