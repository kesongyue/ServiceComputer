package cmd

import (
	"Go-Agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "quit a meeting by title",
	Long: `quit a meeting by title,only as a participator! 
			If you are the sponsor of this meeting, you shoule use the command "cancelMeeting" `,
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			title, _ := cmd.Flags().GetString("title")

			//检查是否是自己发起的会议
			if entity.GetMeetingByTitle(title).Sponsor == entity.Loginuser.Name {
				fmt.Println("You can not quit this meeting because you are the sponsor of this meeting!")
				return
			}
			i := 0
			for i = 0; i < len(entity.GetMeetings()); i++ {
				if entity.GetMeetings()[i].Title == title {
					if !entity.GetMeetings()[i].IsParticipator(entity.Loginuser.Name) {
						fmt.Println("You can not quit this meeting because you don't join in!")
						return
					} else {
						entity.GetMeetings()[i].DeleteParticipator(entity.Loginuser.Name)
					}
					break
				}
			}
			//检查是否存在此会议
			if i >= len(entity.GetMeetings()) {
				fmt.Println("There is no this meeting named " + title)
				return
			}

			var remainMeeting []entity.Meeting
			for _, meet := range entity.GetMeetings() {
				if meet.Title != title || (meet.Title == title && len(meet.Participator) != 0) {
					remainMeeting = append(remainMeeting, meet)
				}
			}

			entity.Meetings = remainMeeting
			entity.WriteJson()
			fmt.Println("quit the meeting named " + title + " Successfully!")
		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.Flags().StringP("title", "t", "", "the Meeting Title")
}
