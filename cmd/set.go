/*
Copyright © 2020 frost.wong <happyhackerwqc@foxmail.com>

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

	"github.com/lovelock/gomemcache/v3/memcache"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set ${host} ${port} ${key} ${value}",
	Short: "set value of ${value} to key ${key}",
	Run: func(cmd *cobra.Command, args []string) {
		mc := memcache.New(fmt.Sprintf("%s:%d", host, port))
		mc.DisableCAS = true
		if err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)}); err != nil {
			return
		}

		fmt.Printf("set ok for key: %s, value: %s\n", key, value)
	},
}

func init() {
	setCmd.Flags().StringVarP(&value, "value", "V", "", "value to set to ${key}")
	setCmd.MarkFlagRequired("value")
	rootCmd.AddCommand(setCmd)
}
