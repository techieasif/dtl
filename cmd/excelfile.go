package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/techieasif/dtl/internal/pkg"
)

func init() {
	rootCmd.AddCommand(excelFile)
}

var excelFile = &cobra.Command{
	Use:   "excelfile",
	Short: "EXCEL to JSON",
	Long:  `Use to convert Excel file to json`,
	Run: func(cmd *cobra.Command, args []string) {
		e := handleExcelPath(args)
		if e != nil {
			log.Fatal(e)
		}
	},
}

func handleExcelPath(args []string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments to process, Excel Path)")
	}

	csvPath := args[0]
	sheetName := args[1]
	pkg.FromEXCELPath(csvPath, sheetName)
	return nil

}
