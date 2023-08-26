package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/techieasif/dtl/internal/pkg"
)

func init() {
	rootCmd.AddCommand(csvFile)
}

var csvFile = &cobra.Command{
	Use:   "csvfile",
	Short: "CSV to JSON",
	Long:  `Use to convert csv file to json`,
	Run: func(cmd *cobra.Command, args []string) {
		e := handleCSVPath(args)
		if e != nil {
			log.Fatal(e)
		}
	},
}

func handleCSVPath(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments to process, CSV Path)")
	}

	csvPath := args[0]

	pkg.FromCSVPath(csvPath)

	return nil

}
