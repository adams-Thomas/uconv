/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adams-Thomas/uconv/lib"

	"github.com/spf13/cobra"
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
	// Check to make sure the converstions file exists
	// If it does the continue as normal, if it doesn't then create it in the lib folder and add base conversions

	_, err := os.Stat("./conversions.txt")
	if err != nil && strings.Contains(err.Error(), "cannot find the file") {
		fmt.Printf("Conversion file not found, adding it...\n\n")
		f, createError := os.Create("./conversions.txt")
		if createError != nil {
			panic(err)
		}

		defaults := "cel,fh,32\nrem,px,16\nmile,km,1.60934\nm,feet,3.28084\n"
		f.WriteString(defaults)
		f.Sync()
	}
}
