package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "A brief description of your command",
	Long:  `create a Meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			sponsor := entity.Loginuser.Name
			title, _ := cmd.Flags().GetString("title")
			participator, _ := cmd.Flags().GetString("participator")
			startTime, _ := cmd.Flags().GetString("starttime")
			endTime, _ := cmd.Flags().GetString("endtime")
			//检查会议是否已创建
			for _, meet := range entity.GetMeetings() {
				if meet.Title == title {
					fmt.Printf("Meeting %s has been create!Create Failed\n", title)
					log.Printf("Meeting %s has been create!Create Failed\n", title)
					return
				}
			}
			participators := strings.Split(participator, ",")
			//检查成员是否是用户或者参与者里包含发起者
			for _, partic := range participators {
				if !entity.IsUser(partic) {
					fmt.Printf("%s is not a user!\nCreate Meeting Failed!", partic)
					return
				}
				if partic == sponsor {
					fmt.Println("Sponsor can not be a pariticipator")
					return
				}
			}
			st := entity.StringToDate(startTime)
			et := entity.StringToDate(endTime)
			if !entity.IsDateValid(st) || !entity.IsDateValid(et) {
				fmt.Println("The startTime or endTime is not valid\nCreate Meeting Failed!")
				return
			}
			if !entity.IsEndBigthanStart(st, et) {
				fmt.Println("The startTime is larger than endTime!\nCreate Meeting Failed!")
				return
			}
			//找出与新创的会议时间冲突的会议
			var conflictMeetings []entity.Meeting
			for _, meet := range entity.GetMeetings() {
				if !(entity.IsEndBigthanStart(et, meet.Start) || entity.IsEndBigthanStart(meet.End, st)) {
					conflictMeetings = append(conflictMeetings, meet)
				}
			}
			//检查每个成员的时间是否合适参加此会议
			for _, partic := range participators {
				for _, meet := range conflictMeetings {
					if meet.IsParticipator(partic) {
						fmt.Println(partic + " has a conflict meeting name " + meet.Title)
						return
					}
				}
			}

			newMeeting := entity.Meeting{
				Sponsor:      sponsor,
				Participator: participators,
				Start:        st,
				End:          et,
				Title:        title,
			}
			entity.AddMeeting(newMeeting)
			entity.WriteJson()
			fmt.Println("Meeting " + title + " create successfully!")

		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	createMeetingCmd.Flags().StringP("title", "t", "", "the Meeting title")
	createMeetingCmd.Flags().StringP("participator", "p", "", "the list of participator in Meeting")
	createMeetingCmd.Flags().StringP("starttime", "s", "", "the Meeting starting time(year-month-day/hour:min)")
	createMeetingCmd.Flags().StringP("endtime", "e", "", "the Meeting end time(year-month-day/hour:mim)")
}
