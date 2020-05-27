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
	gitHubCommand "github.com/oneCoder16/issues-tool/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pubCmd represents the pub command
var addCmd = &cobra.Command{
	Use:     "add [comment data]",
	Aliases: []string{"a"},
	Short:   "add comment",
	Long:    ``,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var params map[string]interface{} = make(map[string]interface{})
		var command string = "add_comment"

		issues_id, err := cmd.Flags().GetInt("issues_id")
		if err != nil {
			fmt.Printf("get issues_id err : %s", err.Error())
			return
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			fmt.Printf("get user err : %s", err.Error())
			return
		}

		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			fmt.Printf("get repo err : %s", err.Error())
			return
		}

		if issues_id != 0 {
			params["issues_id"] = issues_id
		} else {
			params["issues_id"] = viper.GetInt("issues_id")
		}

		if user != "" {
			params["user"] = user
		} else {
			params["user"] = viper.GetString("user")
		}

		if repo != "" {
			params["repo"] = repo
		} else {
			params["repo"] = viper.GetString("repo")
		}

		params["content"] = args[0]

		if err := gitHubCommand.NewBroker().Execute(command, params); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	addCmd.Flags().StringP("user", "u", "", "github user")
	addCmd.Flags().StringP("repo", "r", "", "github user repo")
	addCmd.Flags().IntP("issues_id", "i", 0, "issues id")

	rootCmd.AddCommand(addCmd)
}
