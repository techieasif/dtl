/*
Copyright © 2023 Mohammad Asif <techieasif@gmail.com>

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
	"github.com/techieasif/dtl/internal/pkg"
	"log"

	"github.com/spf13/cobra"
)

// base64decodeCmd represents the base64decode command
var base64decodeCmd = &cobra.Command{
	Use:   "base64decode",
	Short: "Decode an encoded string with base 64",
	Long:  `Decode your data, with ease in your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handleBase64Decode(args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func handleBase64Decode(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments to process")
	}
	arg := args[0]

	result, err := pkg.Base64Decode(arg)

	if err != nil {
		return err
	}

	fmt.Println("decoded: ", *result)
	return nil

}
func init() {
	rootCmd.AddCommand(base64decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
