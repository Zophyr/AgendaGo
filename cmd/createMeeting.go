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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting -t [meeting title] -p [participators] -s [start time] -e [end time]",
	Short: "create a meeting",
	Long:  `e.g. createMeeting -t MixShow -p a b c -s 2018-11-11/10:00 -e 2018-11-11/14:00`,

	Run: func(cmd *cobra.Command, args []string) {

		theTitle, _ := cmd.Flags().GetString("title")
		theParticipators, _ := cmd.Flags().GetStringArray("participator")
		theStart, _ := cmd.Flags().GetString("startTime")
		theEnd, _ := cmd.Flags().GetString("endTime")
		err := service.AddMeetingToCurrSession(theTitle, theParticipators, theStart, theEnd)

		if err == nil {
			fmt.Println("Add meeting: ", theTitle, " successfully!")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}

	},
}

func init() {

	rootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("title", "t", "", "title of the meeting")
	createMeetingCmd.Flags().StringArrayP("participator", "p", nil, "participators of the meeting")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "start time of the meeting")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "end time of the meeting")
}
