/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/lovelock/mclient/core"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set ${host} ${port} ${key} ${value}",
	Short: "set value of ${value} to key ${key}",
	Run: func(cmd *cobra.Command, args []string) {
		core.Set(key, value, host, port)
	},
}

func init() {
	setCmd.Flags().StringVarP(&value, "value", "V", "", "value to set to ${key}")
	setCmd.MarkFlagRequired("value")
	rootCmd.AddCommand(setCmd)
}
