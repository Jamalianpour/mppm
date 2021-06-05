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
	"os"

	"github.com/jamalianpour/mppm/model"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Set a name for your project")
		}

		exist, _ := pathExists("./mppm/project")
		if !exist {
			os.MkdirAll("./mppm/project", os.ModePerm)
		}
		project := model.ProjectModel{
			Name: args[0],
		}

		jsonString, _ := json.Marshal(project)

		err := ioutil.WriteFile("./mppm/project/"+args[0]+".json", jsonString, os.ModePerm)
		if err != nil {
			return errors.New(err.Error())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0] + " created successfully 🎉\n ")
		fmt.Println("Now add command to project with 'mppm add [ProjectName] --command=***'")
		fmt.Println("Use \"mppm add --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
