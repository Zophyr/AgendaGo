// Copyright © 2018 Zophyr <the-zephyr@hotmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		fmt.Println("register called by " + username + " password " + password + " email " + email + " phone " + phone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	
	registerCmd.Flags().StringP("username", "u", "", "用户名")
	registerCmd.Flags().StringP("password", "p", "", "密码")
	registerCmd.Flags().StringP("email", "m", "","邮箱")
	registerCmd.Flags().StringP("phone","t", "","电话")
}
