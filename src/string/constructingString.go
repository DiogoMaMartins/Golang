package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0x66, 0xc3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}
