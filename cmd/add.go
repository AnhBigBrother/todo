package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long:  "Add new todo, ex: todo add <todo_title>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, todo := range args {
			todos.Add(todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
