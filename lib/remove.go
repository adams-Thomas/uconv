package lib

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Remove(args []string) {
	filePath := viper.GetString("filePath")
	file, err := ioutil.ReadFile(filePath)
	Check(err)

	var conversion string
	var index int
	fileSplit := strings.Split(string(file), "\n")
	for i, current := range fileSplit {
		split := strings.Split(current, ",")
		check := fmt.Sprintf("%s%s", split[0], split[1])
		if check == args[0] {
			conversion = current
			index = i
			break
		}
	}

	newSplit := fileSplit[:index]
	newSplit = append(newSplit, fileSplit[index+1:]...)

	edit, err := os.Create(filePath)
	Check(err)
	defer edit.Close()
	edit.WriteString(strings.Join(newSplit, "\n"))
	edit.Sync()

	// scanner := bufio.NewScanner(file)
	// var conversion string
	// for scanner.Scan() {
	// 	current := scanner.Text()
	// 	split := strings.Split(current, ",")
	// 	check := fmt.Sprintf("%s%s", split[0], split[1])
	// 	if check == args[0] {
	// 		conversion = current
	// 		break
	// 	}
	// }

	// if conversion == "" {
	// 	fmt.Printf("%s was not found, please try again", args[0])
	// 	return
	// }

	// fmt.Println(scanner.Bytes())

	fmt.Printf("%s has been removed", conversion)
}
