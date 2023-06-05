/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adams-Thomas/uconv/lib"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uconv",
	Short: "A basic CLI for unit conversions",
	Args:  cobra.ExactArgs(2),
	Long: `
uconv is a little CLI app for unit conversions without needing to Google.
Features:
	- Base Usage
	- Add: Add your own conversions [uconv Add -h/--help]
	- Delete: Delete conversions [uconv Delete -h/--help]
	- List: List the current conversions [uconv List -h/--help]

Basic usage example:
		
To convert from rem to px run the following command:
	uconv rempx 8
		
The same command can be run "backwards"
	uconv pxrem 32
This will use the same conversion data
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		lib.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	/**
	Check if the uconv folder and text file exits
	This code is starting to become messy, should really find a way to clean it up.

	First checks to make sure the uconv folder exists in app dir
	If it does not then make it, then make the conversions file because it will be missing if
	uconv dir is missing

	If the uconv dir is there, double check the conversions file exists because it is possible
	the conversions file can be missing but the dir still be there.
	*/

	dirPath := fmt.Sprintf("%s\\uconv", xdg.DataHome)
	filePath := fmt.Sprintf("%s\\uconv\\conversions.txt", xdg.DataHome)
	_, err := os.Stat(dirPath)

	if err != nil && strings.Contains(err.Error(), "cannot find the file") {
		fmt.Println("uconv dir not found, adding it")
		folderErr := os.Mkdir(dirPath, os.ModePerm)

		if folderErr != nil {
			panic(err)
		}
	}

	_, fileErr := os.Stat(filePath)
	if fileErr != nil && strings.Contains(fileErr.Error(), "cannot find the file") {
		fmt.Println("Base conversions data not found, adding")
		f, createError := os.Create(filePath)
		if createError != nil {
			panic(err)
		}

		defaults := "cel,fh,32\nrem,px,16\nmile,km,1.60934\nm,feet,3.28084\n"
		f.WriteString(defaults)
		f.Sync()
	}

	viper.SetDefault("dirPath", dirPath)
	viper.SetDefault("filePath", filePath)
}
