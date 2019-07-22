package main

import (
	"fmt"
)

func main() {
	var input float32

	fmt.Print("Enter a floating number: ")
	fmt.Scanln(&input)

	//using type conversion to convert to int
	intNum := int(input)

	fmt.Printf("This is the truncated integer: %d\n", intNum)
}
