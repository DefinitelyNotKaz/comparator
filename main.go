package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Options struct {
	File    string
	Palette string
	Verbose bool
	Output  string
}

var options Options

var rootCmd = &cobra.Command{
	Use:   "comparator",
	Short: "Compare pixel art images to palette files",
	Long:  "Compare pixel art images to palette files",
	Run: func(cmd *cobra.Command, args []string) {
		err := Compare(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&options.File, "input", "i", "", "Input file (required)")
	rootCmd.Flags().StringVarP(&options.Palette, "palette", "p", "", "Palette to compare against (required)")
	rootCmd.Flags().StringVarP(&options.Output, "output", "o", "result.png", "File to save the result (default: 'result.png').")
	rootCmd.Flags().BoolVar(&options.Verbose, "verbose", false, "Output extra information")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("palette")
}

func main() {
	Execute()
}
