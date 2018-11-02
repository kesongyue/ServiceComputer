package cmd

import (
	"Go-Agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

// queryUserCmd represents the queryUser command
var queryUserCmd = &cobra.Command{
	Use:   "queryUser",
	Short: "A brief description of your command",
	Long:  `query all users`,
	Run: func(cmd *cobra.Command, args []string) {
		if entity.IsLogin() {
			for i := 0; i < len(entity.Users); i++ {
				fmt.Printf("%s  %s  %s\n", entity.Users[i].Name, entity.Users[i].Email, entity.Users[i].Phone)
			}
		} else {
			fmt.Println("You must login first")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
