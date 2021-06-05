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
	"strings"

	"github.com/jamalianpour/mppm/model"
)

var argsErr error

func readFile(name string) (model.ProjectModel, error) {
	data, err := ioutil.ReadFile("./mppm/project/" + name + ".json")

	var projectJson model.ProjectModel

	err = json.Unmarshal(data, &projectJson)
	return projectJson, err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func projectExists(name string) error {
	_, argsErr = os.Stat("./mppm/project/" + name + ".json")
	if os.IsNotExist(argsErr) {
		err := errors.New("Error: The " + name + " dose not found!!! ⛔\nFirst cearte a project then add command. You can create a project: mppm create " + name)
		return err
	}
	return nil
}

func removeProject(name string) error {
	fmt.Println("Are you sure? ")
	var agreed string
	fmt.Scan(&agreed)
	agreed = strings.ToLower(agreed)
	if agreed == "y" || agreed == "yes" {
		return os.Remove("./mppm/project/" + name + ".json")
	}
	return nil
}

func RemoveIndex(s []model.Commands, index int) []model.Commands {
	return append(s[:index], s[index+1:]...)
}
