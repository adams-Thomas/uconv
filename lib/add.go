package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensure 3rd arg is a number/value/float
// First checks to make sure the conversion doesn't exist
// If it doesn't, then add it
func Add(args []string) {
	_, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Printf("Error: %s is not a valid number.", args[2])
		return
	}

	dat, fileErr := os.OpenFile("./conversions.txt", os.O_RDWR, 0644)
	if fileErr != nil {
		panic(err)
	}
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		conversion := scanner.Text()
		check := fmt.Sprintf("%s,%s", args[0], args[1])
		if strings.Contains(conversion, check) {
			fmt.Printf("The conversion you are trying to add already exists. \n")
			fmt.Printf("\nYou are trying to add: \n  %s,%s,%s", args[0], args[1], args[2])
			fmt.Printf("\n\nThe conversion already exists with: \n  %s", conversion)
			fmt.Printf("\n\nTo change or edit the conversion, delete it and add it back\n")
			return
		}
	}

	// writer := bufio.NewWriter(dat)
	newConversion := fmt.Sprintf("%s,%s,%s\n", args[0], args[1], args[2])
	_, writeErr := dat.WriteString(newConversion)

	if writeErr != nil {
		fmt.Println("Error: Unable to write to file. Please try again")
		panic(err)
	}
	dat.Sync()

	fmt.Printf("New conversion added: \n\n")
	fmt.Printf("  %s,%s,%s\n", args[0], args[1], args[2])
}
