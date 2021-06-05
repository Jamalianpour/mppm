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

	"github.com/jamalianpour/mppm/model"
	"github.com/spf13/cobra"
)

var projectName, command, path string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add command to your project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Select a project first and then add command to that with --command(-c) and --path(-p)")
		}
		projectName = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if argsErr == nil {
			projectJson, err := readFile(projectName)
			if err != nil {
				fmt.Print(err)
			}

			flagPath, _ := cmd.Flags().GetString("path")
			flagCommand, _ := cmd.Flags().GetString("command")

			newCommand := model.Commands{
				Path:    flagPath,
				Command: flagCommand,
			}

			projectJson.Commands = append(projectJson.Commands, newCommand)

			jsonString, _ := json.Marshal(projectJson)

			err = ioutil.WriteFile("./mppm/project/"+args[0]+".json", jsonString, os.ModePerm)
			if err != nil {
				fmt.Println("error:", err)
			}

			fmt.Println("New command add to project successfully 🎉")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		err := projectExists(projectName)
		if err != nil {
			log.Fatal(err)
		}
		// _, argsErr = os.Stat("./mppm/project/" + projectName + ".json")
		// if os.IsNotExist(argsErr) {
		// 	err := errors.New("Error: The " + projectName + " dose not found!!! ⛔")
		// 	fmt.Println(err)
		// 	fmt.Println("First cearte a project then add command. You can create a project: mppm create " + projectName)
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("path", "p", "", "Command execute at this path")

	addCmd.Flags().StringP("command", "c", "", "Command to execute IDE, tools or etc")

	addCmd.MarkFlagRequired("path")
	addCmd.MarkFlagRequired("command")
}
