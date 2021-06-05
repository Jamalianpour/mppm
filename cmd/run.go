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
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run and execute project commands",
	Long:    "Run and execute project commands Like this: mppm run [projectName]",
	Example: "mppm run myProject",
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
				os := runtime.GOOS
				var app string

				switch os {
				case "windows":
					app = "powershell"
				case "darwin":
					app = "terminal"
				case "linux":
					app = "bash"
				default:
					log.Fatal("⛔ Sorry, your os dose not supported ⛔")
					return
				}

				for i := 0; i < len(proj.Commands); i++ {
					fmt.Println("Run Command " + strconv.Itoa(i) + ": " + proj.Commands[i].Command)
					cmd := exec.Command(app, "-c", proj.Commands[i].Command)
					cmd.Dir = proj.Commands[i].Path
					stdout, err := cmd.Output()

					if err != nil {
						fmt.Println("⛔ " + err.Error() + " ⛔\n")
						// return
					} else {
						// Print the output
						fmt.Println(string(stdout))
					}

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
	rootCmd.AddCommand(runCmd)
}
