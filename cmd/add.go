package cmd

import (
	"fmt"
	"strings"

	"github.com/patrickblackjr/taskctl/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Add \"%s\" to your todo list", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
