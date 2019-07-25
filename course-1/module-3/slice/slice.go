package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sli := make([]int, 0, 3)

	for {
		fmt.Print("Please enter an integer (enter 'X' to exit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		intNew, _ := strconv.Atoi(scanner.Text())

		// break to exit for-loop
		if strings.ToUpper(scanner.Text()) == "X" {
			fmt.Println("\n'X' entered, exiting program. Goodbye!")
			break
		}
		//add integer to slice
		sli = append(sli, intNew)
		//sorting
		sort.Sort(sort.IntSlice(sli))
		//printing the slice
		fmt.Println("This is the slice", sli)
	}
}
