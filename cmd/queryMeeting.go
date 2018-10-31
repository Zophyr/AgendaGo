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

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "query the meeting by its title",
	Long: `input is the title,then return the class of the meeting which include
	 date and participator`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		results, err := service.QueryMeeting(startTime, endTime)
		if err == nil {
			if len(results) == 0 {
				fmt.Printf("No meeting can be found")
				return
			}
			for i, meeting := range results {
				fmt.Printf("No. %d \nTitle: %s Speecher: %s\n", i, meeting.Title, meeting.Sponsor)
				fmt.Printf("Participators:\n")
				for _, participator := range meeting.Participators {
					fmt.Printf("%s\n", participator)
				}
				fmt.Printf("Start Time: %s  End Time: %s\n\n", meeting.StartTime, meeting.EndTime)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "start time of the query interval")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "end time of the query interval")
}
