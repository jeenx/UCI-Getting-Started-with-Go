package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter First name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter address: ")
	scanner.Scan()
	addr := scanner.Text()

	var m = map[string]string{
		name: addr}

	jsonString, _ := json.Marshal(m)

	fmt.Println(string(jsonString))
}
