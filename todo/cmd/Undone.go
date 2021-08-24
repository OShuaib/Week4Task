/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	source "TodoList/sourceCode"
	"strconv"

	"github.com/spf13/cobra"
)

// UndoneCmd represents the Undone command
var UndoneCmd = &cobra.Command{
	Use:   "Undone",
	Short: "Undone your task on the list ",
	Long: `Gives you the power to uncheck your incomplete task 
	
	and to help reverse the mistake of marking an undone task`,
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		source.Task.Undone(index)
	},
}

func init() {
	rootCmd.AddCommand(UndoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UndoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UndoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
