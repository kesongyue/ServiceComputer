package cmd

import (
	"Go-Agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "delete a meeting by title",
	Long:  "delete a meeting by title, only you can delete the meeting sponsored by you",
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			title, _ := cmd.Flags().GetString("title")

			//抽取剩下的meeting
			var remainMeeting []entity.Meeting
			for i := 0; i < len(entity.GetMeetings()); i++ {
				if entity.GetMeetings()[i].Title == title {
					if entity.GetMeetings()[i].Sponsor != entity.Loginuser.Name {
						fmt.Println("You can not cancel a meeting that is not sponsored by you!")
						return
					}
				} else {
					remainMeeting = append(remainMeeting, entity.GetMeetings()[i])
				}
			}
			//检查是否存在此会议
			if len(remainMeeting) == len(entity.GetMeetings()) {
				fmt.Println("There is no this meeting named " + title)
				return
			}

			entity.Meetings = remainMeeting
			entity.WriteJson()
			fmt.Println("Cancel the meeting named " + title + " Successfully!")
		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelMeetingCmd)
	cancelMeetingCmd.Flags().StringP("title", "t", "", "the Meeting Title")
}
