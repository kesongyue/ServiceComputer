package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addMeetingCmd represents the addMeeting command
var addMeetingMemberCmd = &cobra.Command{
	Use:   "addMeetingMember",
	Short: "add MeetingMember by title",
	Long:  "add MeetingMember by title,usage:addMeetingMember -t xxx -p xxx,xxx,...",
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			title, _ := cmd.Flags().GetString("title")
			participator, _ := cmd.Flags().GetString("participator")
			// 检查改会议的发起者是不是已登录的用户
			meeting := entity.GetMeetingByTitle(title)
			if meeting.Sponsor != entity.Loginuser.Name {
				fmt.Println("Sorry! The Meeting " + title + " is not sponsored by you, you have no access to add a participator")
				return
			}
			//检查每个参与者是否是用户
			participators := strings.Split(participator, ",")
			for _, partic := range participators {
				if !entity.IsUser(partic) {
					fmt.Println("There is no this user named " + partic)
					return
				}
			}
			//找出与该会议时间冲突的会议
			var conflictMeetings []entity.Meeting
			for _, meet := range entity.GetMeetings() {
				if !(entity.IsEndBigthanStart(meeting.End, meet.Start) || entity.IsEndBigthanStart(meet.End, meeting.Start)) {
					conflictMeetings = append(conflictMeetings, meet)
				}
			}
			//检查是否有参与者在冲突会议里
			for _, partic := range participators {
				for _, meet := range conflictMeetings {
					if meet.IsParticipator(partic) {
						fmt.Println(partic + " has a conflict meeting name " + meet.Title)
						return
					}
				}
			}

			for _, partic := range participators {
				meeting.Participator = append(meeting.Participator, partic)
			}
			// 要用下标访问才能改变数组的内容
			for i := 0; i < len(entity.GetMeetings()); i++ {
				if entity.GetMeetings()[i].Title == meeting.Title {
					entity.GetMeetings()[i] = meeting
				}
			}

			entity.WriteJson()
			fmt.Println("Add Meeting Participator  successfully!")

		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addMeetingMemberCmd)
	addMeetingMemberCmd.Flags().StringP("title", "t", "", "the Meeting title")
	addMeetingMemberCmd.Flags().StringP("participator", "p", "", "the Meeting participator")
}
