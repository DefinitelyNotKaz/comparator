package main

import (
	"os"

	"github.com/spf13/cobra"
)

var file string
var template string
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "comparator",
	Short: "Compare pixel art images to palette files",
	Long:  "Compare pixel art images to palette files",
	Run: func(cmd *cobra.Command, args []string) {
		compare(file, template, verbose)
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define the flags and their default values
	rootCmd.Flags().StringVarP(&file, "input", "i", "", "Input file (required)")
	rootCmd.Flags().StringVarP(&template, "palette", "p", "", "Palette to compare against (required)")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Output extra information")

	// Mark the input flag as required
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("palette")
}

func main() {
	Execute()
}
