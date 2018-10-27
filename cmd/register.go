// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"AgendaGo/service"
	"fmt"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Used to register account",
	Long:  `User can register itself by username, password, email and phone.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		err := service.RegisterUser(username, password, email, phone)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(username + " registers successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	registerCmd.Flags().StringP("username", "u", "", "/")
	registerCmd.Flags().StringP("password", "p", "", "/")
	registerCmd.Flags().StringP("email", "m", "", "/")
	registerCmd.Flags().StringP("phone", "t", "", "/")
}
