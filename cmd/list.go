/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"uconv/lib"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Command to list all available conversions",
	Long: `
uconv list <ARGS>
This is used to list and show all the available conversions for uconv.

ARGS:
	- <FORMATTING> --> Optional argument

<FORMATTING> has 4 possible values:
	1. Left empty [Defaults to format]
			uconv list --> 1. From 1cel to 32fh
	2. format
			uconv list format --> 1. From 1cel to 32fh
	3. plain
			uconv list plain --> cel,fh,32
	4. usages
			uconv list usages --> celfh
	`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.List(args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// f -> default Formatted like how I have it
	// p -> plain from the text file
	// c -> just the available conversions can be reversed
}
