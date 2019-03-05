package main

import "fmt"

func main () {
	calculateBill(10,5)
}


func calculateBill(price int, no int) int {
	var totalPrice = price * no
	fmt.Println("total",totalPrice)
	return totalPrice
}

