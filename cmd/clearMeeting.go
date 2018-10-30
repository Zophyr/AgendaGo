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
	"os"

	"github.com/spf13/cobra"
)

// clearMeetingCmd represents the clearMeeting command
var clearMeetingCmd = &cobra.Command{
	Use:   "clearMeeting",
	Short: "clear all the meeting",
	Long:  `by running this clear all the register meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.DeleteAllMeeting()
		if err == nil {
			fmt.Println("Cleared all the meetings")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
