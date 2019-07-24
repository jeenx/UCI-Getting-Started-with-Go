package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	// Conver input to lower case
	finalString := strings.ToLower(input)

	if finalString[:1] == "i" && finalString[len(finalString)-1:] == "n" && strings.Contains(finalString, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
