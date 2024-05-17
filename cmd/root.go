package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quake",
	Short: "Quake CLI application",
	Long:  `A command line application to process Quake log files and generate reports in JSON format.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
