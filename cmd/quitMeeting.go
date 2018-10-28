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
	"fmt"

	"github.com/spf13/cobra"
)

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "help the current user to quit the correctsponding meeting by its title",
	Long: `using the input of title,adding the other which is the current userName then delete 
	the participator of the meeting,after that check the correctsponding meeting if the participator
	is all clear then delete the meeting`,
	Run: 
		func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		err := service.DeleteFromMeeting(title) 
		if service.queryMeeting(title).getParticipator().empty()==true{
			service.DeleteMeetingByTitle("title")
		}
		if err == nil {
			fmt.Printf("Quited the meeting %s\n", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	quitMeetingCmd.Flags().StringP("title", "t", "", "/")
}
