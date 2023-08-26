package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dtl",
		Short: "Dev utils ",
		Long:  `DTL aka detel is used for development related utils`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
