package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long:  `Login`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username == "" || password == "" {
			fmt.Println("Please check your input")
			log.Println("Please check your input")
			return
		}

		var users = entity.GetUsers()
		var flag = false
		for i := 0; i < len(users); i++ {
			if users[i].Name == username && users[i].Password == password {
				entity.Loginuser = users[i]
				flag = true
			}
		}

		if flag == false {
			fmt.Println(username + " login fail")
			log.Println(username + " login fail")
		} else {
			fmt.Println(username + " login successful")
			log.Println(username + " login successful")
			entity.WriteLoginJson()
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "the user name")
	loginCmd.Flags().StringP("password", "p", "", "the password")
}
