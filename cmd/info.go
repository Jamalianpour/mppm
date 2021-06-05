/*
Copyright © 2021 Mohammad Jamalianpour <jamalianm21@yahoo.com>

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
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show all detail of a project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Select a project! ⛔")
		}
		projectName = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if argsErr == nil {
			proj, err := readFile(projectName)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Name: " + proj.Name)
				for i := 0; i < len(proj.Commands); i++ {
					fmt.Println("---------------", "Command", i, "---------------")
					fmt.Println("Path: " + proj.Commands[i].Path)
					fmt.Println("Command: " + proj.Commands[i].Command)
				}
			}
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		err := projectExists(projectName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
