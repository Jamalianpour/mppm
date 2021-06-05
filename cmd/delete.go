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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a command of project",
	Long:    "Delete a command of project Like this: mppm delete [projectName] -c [commandIndex]",
	Example: "mppm delete myProject -c 0",
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
				flagCommand, _ := cmd.Flags().GetInt("command")
				if flagCommand < 0 || int(flagCommand) > len(proj.Commands) {
					log.Fatal("Error: wrong command index ⛔ you can see command(s) of project with: mppm info " + projectName)
				}
				for i := 0; i < len(proj.Commands); i++ {
					fmt.Println("---------------", "Command", i, "---------------")
					fmt.Println("Path: " + proj.Commands[i].Path)
					fmt.Println("Command: " + proj.Commands[i].Command)
				}

				fmt.Println("\nDo you want to delete command", flagCommand, "? ")
				var agreed string
				fmt.Scan(&agreed)
				agreed = strings.ToLower(agreed)
				if agreed == "y" || agreed == "yes" {
					proj.Commands = RemoveIndex(proj.Commands, int(flagCommand))
					jsonString, _ := json.Marshal(proj)
					err = ioutil.WriteFile("./mppm/project/"+args[0]+".json", jsonString, os.ModePerm)
					if err != nil {
						log.Fatal("Error:", err)
					}
					fmt.Println("Deleted successfully")
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
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("command", "c", 0, "index of command")

	deleteCmd.MarkFlagRequired("command")
}
