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
	"fmt"

	"github.com/lovelock/gomemcache/memcache"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get ${host} ${port} ${key}",
	Short: "get value of ${key} from memcached server ${host}:${port}",
	Run: func(cmd *cobra.Command, args []string) {
		mc := memcache.New(fmt.Sprintf("%s:%d", host, port))
        mc.DisableCAS = true

		value, err := mc.Get(key)
		if err != nil {
			fmt.Printf("error occurred: %s\n", err)
		}

		fmt.Println(string(value.Value))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
