package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type conversion struct {
	from    string
	to      string
	value   float64
	reverse bool
}

func (c *conversion) Calc(num float64) float64 {
	if c.reverse == true {
		return num / c.value
	}

	return c.value * num
}

func ConversionFromSlice(data []string) *conversion {
	value, err := strconv.ParseFloat(data[2], 64)
	if err != nil {
		fmt.Printf(`
		%s is not a valid number. 
		Please remove the conversion and re-add it with the correct value.
		`, data[2])
		return nil
	}

	reverse := false
	if len(data) == 4 && data[3] == "reverse" {
		reverse = true
	}

	return &conversion{
		from:    data[0],
		to:      data[1],
		value:   value,
		reverse: reverse,
	}
}

func Run(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println(`
		To few arguments, please enter conversion type and value.
		Type uconv --help for more
		`)
		return
	}

	conv := args[0]
	num, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		fmt.Println(`
		Please enter a valid number for the conversion value.
		Type uconv --help for more
		`)
		return
	}

	conversionSlice := SearchFile(conv)

	if len(conversionSlice) == 0 {
		fmt.Printf(`
		The conversion, %s, entered was not found.
		Please make sure entered the conversion correctly, or
		the conversion is present in the available list.
		To check enter: uconv List
		`, conv)
		return
	}

	convObj := ConversionFromSlice(conversionSlice)

	if convObj == nil {
		return
	}

	result := convObj.Calc(num)
	fmt.Println(
		result, convObj.to,
	)
}

func SearchFile(conv string) []string {
	filePath := viper.GetString("filePath")
	dat, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	var conversion []string

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")

		if split[0]+split[1] == conv {
			conversion = split
			break
		}

		if split[1]+split[0] == conv {
			conversion = split
			conversion = append(conversion, "reverse")
			break
		}
	}

	return conversion
}
