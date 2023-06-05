/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/adams-Thomas/uconv/lib"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <FROM> <TO> <VALUE>",
	Args:    cobra.ExactArgs(3),
	Example: "  uconv add rem px 16\n\nWhere 1 rem is equal to 16 px\nTo use the conversion then run:\n  uconv rempx <value>",
	// Args:    cobra.ExactArgs(3),
	Short: "Add a new conversion cli",
	Long: `
The add command allows you to add conversions to the list that do not come by default.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
