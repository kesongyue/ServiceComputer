package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `Delete a user from agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			var NewUsers []entity.User
			for i := 0; i < len(entity.Users); i++ {
				if entity.Users[i].Name != entity.Loginuser.Name && entity.Users[i].Password != entity.Loginuser.Password {
					NewUsers = append(NewUsers, entity.Users[i])
				}
			}
			entity.Users = NewUsers
			entity.WriteJson()
			f, _ := os.OpenFile(entity.ULoginFilepath, os.O_WRONLY|os.O_TRUNC, 0600)
			defer f.Close()
			fmt.Println("Delete successfully")
		} else {
			fmt.Println("You must login first")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
