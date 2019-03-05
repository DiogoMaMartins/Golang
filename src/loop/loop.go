package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
	  if i > 5 {
	    break
	  }
	  fmt.Printf("%d", i)
	}
	fmt.Println("\nline after for loop")
	continuee()
}


func continuee() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("%d", i)
	}
}


func outer(){
  outer:
		for i := 0; i < 3; i++ {
			for j:= 1; j < 4; j++ {
				fmt.Println("i = %d, j = %d\n",i,j)
				if i == j {
				break outer
				}
			}
		}
}
