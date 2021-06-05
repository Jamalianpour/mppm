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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all project in mppm",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("./mppm/project")
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}
		fmt.Println(" ____________________________________")
		fmt.Println("|List of all project created in mppm:|")
		for i, f := range files {
			output := fmt.Sprintf("|%d%s%s", i+1, ". ", strings.TrimSuffix(f.Name(), ".json"))
			space := 37 - len(output)
			if space > 0 {
				for j := 0; j < space; j++ {
					output = output + " "
				}
			}
			output = output + "|"
			// fmt.Println("|____________________________________|")
			fmt.Println(output)
		}
		fmt.Println("|____________________________________|")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
