/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/techieasif/dtl/internal/pkg"
	"log"
)

// base64encodeCmd represents the base64encode command
var base64encodeCmd = &cobra.Command{
	Use:   "base64encode",
	Short: "Encode String with base 64",
	Long:  `Encode your data, with ease in your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handleBase64Encode(args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func handleBase64Encode(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments to process, Excel Path)")
	}
	arg := args[0]

	result, err := pkg.Base64Encode(arg)

	if err != nil {
		return err
	}

	fmt.Println("encoded: ", *result)
	return nil

}

func init() {
	rootCmd.AddCommand(base64encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
