package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FORMATTING options
const (
	FORMAT = "format"
	PLAIN  = "plain"
	USAGES = "usages"
)

func List(args []string) {
	if len(args) > 1 {
		fmt.Println(`
		To many arguments supplied, only one required. see uconv list -h for more`)
		return
	}

	option := FORMAT
	if len(args) != 0 {
		option = args[0]
	}

	switch option {
	default:
	case FORMAT:
		// {index}. From 1 {from} to {value} {to}
		ListFile(FORMAT)
	case PLAIN:
		// Straight from text file
		ListFile(PLAIN)
	case USAGES:
		// {from}{to}
		ListFile(USAGES)
	}
}

func ListFile(formatting string) {
	dat, err := os.Open("./conversions.txt")

	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)

	index := 1
	for scanner.Scan() {
		conversion := scanner.Text()
		conversionSplit := strings.Split(conversion, ",")

		switch formatting {
		case FORMAT:
			fmt.Printf("%d. From 1 %s to %s %s\n",
				index, conversionSplit[0], conversionSplit[2], conversionSplit[1])

		case PLAIN:
			fmt.Println(conversion)

		case USAGES:
			fmt.Printf("%s%s\n", conversionSplit[0], conversionSplit[1])
		}
		index++
	}
}
