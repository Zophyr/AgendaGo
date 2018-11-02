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

// add participators to a meeting
var addCmd = &cobra.Command{
	Use:   "add participator",
	Short: "add participators to a meeting",
	Long:  `add some existed users to a particular meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participator")
		err := service.AddParticipatorToMeeting(title, participators)
		if err == nil {
			fmt.Printf("Added %s to the meeting %s\n", participators, title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("participator", "p", nil, "the new participator of the meeting")
	addCmd.Flags().StringP("title", "t", "", "the title of the meeting")
}
