package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long:  `Register a user`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("userName")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		if username == "" || password == "" || email == "" || phone == "" {
			fmt.Println("Please check your input")
			log.Println("Please check your input")
			return
		}

		newUser := entity.User{
			Name:     username,
			Password: password,
			Email:    email,
			Phone:    phone,
		}
		entity.AddUser(newUser)
		entity.WriteJson()
		fmt.Println(username + " registers successful")
		log.Println("Please check your input")

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("userName", "u", "", "the user name")
	registerCmd.Flags().StringP("password", "p", "", "the password")
	registerCmd.Flags().StringP("email", "e", "", "the email")
	registerCmd.Flags().StringP("phone", "t", "", "the phone")
}
