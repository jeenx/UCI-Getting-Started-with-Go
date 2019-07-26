package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Person struct with field names
type Person struct {
	fname string
	lname string
}

func main() {
	sli := make([]Person, 0, 20) // created slice with struct

	fmt.Print("Enter file name: ")
	reader := bufio.NewReader(os.Stdin)
	file, _ := reader.ReadString('\n')

	file = strings.Trim(file, "\n")
	openFile, err := os.Open(file) //open file with read access

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(openFile)
	// for loop to add values to slice
	for fileScanner.Scan() {

		fileString := strings.Split(fileScanner.Text(), " ") //split string by white space
		var data Person

		data.fname = fileString[0]
		data.lname = fileString[1]

		sli = append(sli, data)
	}

	//print values in slice
	for _, names := range sli {
		fmt.Println(names)
	}

	openFile.Close()
}
