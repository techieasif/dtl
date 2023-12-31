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
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate an UUID instantly in your terminal",
	Long:  `Want an UUID fir dtl uuid, an id will be created right there.`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}
		uuidStr := string(newUUID)

		fmt.Println("Generated UUID: ", uuidStr)
		fmt.Print("UUID without hyphen(-): ", strings.Replace(uuidStr, "-", "", -1))

	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
