package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todos",
	Long:  "Add new todos (allow multiple), ex: todo add <todo_title1> <todo_title2>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, todo := range args {
			todos.Add(todo)
		}
		todos.PrintAll()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
