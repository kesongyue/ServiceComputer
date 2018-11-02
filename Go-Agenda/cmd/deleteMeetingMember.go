package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// deleteMeetingMemberCmd represents the deleteMeetingMember command
var deleteMeetingMemberCmd = &cobra.Command{
	Use:   "deleteMeetingMember",
	Short: "delete a or a list of MeetingMember in a specifed meeting",
	Long:  "delete a or a list of MeetingMember in a specifed meeting",
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			title, _ := cmd.Flags().GetString("title")
			participator, _ := cmd.Flags().GetString("participator")

			// 检查改会议的发起者是不是已登录的用户
			meeting := entity.GetMeetingByTitle(title)
			if meeting.Sponsor != entity.Loginuser.Name {
				fmt.Println("Sorry! The Meeting " + title + " is not sponsored by you, you have no access to delete a participator")
				return
			}

			//检查发起者是否在参与者列表里，每个参与者是否是用户,检查每个参与者是否在这个会议里
			participators := strings.Split(participator, ",")
			for _, partic := range participators {
				if partic == entity.Loginuser.Name {
					fmt.Println("Sorry! You can not delete yourself as a pariticipator!")
					return
				}
				if !entity.IsUser(partic) {
					fmt.Println("There is no this user named " + partic)
					return
				}
				if !meeting.IsParticipator(partic) {
					fmt.Println(partic + " is not in this meeting named " + title)
					return
				}
			}

			var remainPartic []string
			isRemain := true
			for _, member := range meeting.Participator {
				for _, partic := range participators {
					if member == partic {
						isRemain = false
						break
					}
				}
				if isRemain {
					remainPartic = append(remainPartic, member)
				}
				isRemain = true
			}

			// 要用下标访问才能改变数组的内容
			for i := 0; i < len(entity.GetMeetings()); i++ {
				if entity.GetMeetings()[i].Title == meeting.Title {
					entity.GetMeetings()[i].Participator = remainPartic
				}
			}
			//如果删除后的参与者人数为0就把该会议删掉
			if len(remainPartic) == 0 {
				var remainMeeting []entity.Meeting
				for i := 0; i < len(entity.GetMeetings()); i++ {
					if entity.GetMeetings()[i].Title != meeting.Title {
						remainMeeting = append(remainMeeting, entity.GetMeetings()[i])
					}
				}
				fmt.Println("The number of participator in this meeting is zero!This meeting will be deleted!")
				entity.Meetings = remainMeeting
			}

			entity.WriteJson()
			fmt.Println("Delete Meeting Participator successfully!")
		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingMemberCmd)
	deleteMeetingMemberCmd.Flags().StringP("title", "t", "", "the Meeting title")
	deleteMeetingMemberCmd.Flags().StringP("participator", "p", "", "the Meeting participator")
}
