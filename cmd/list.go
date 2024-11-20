package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completed *bool
var uncompleted *bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  "List all todos as a table",
	Run: func(cmd *cobra.Command, args []string) {
		if *completed && !*uncompleted {
			fmt.Println("Completed tasks: ")
			todos.PrintCompleted()
		} else if !*completed && *uncompleted {
			fmt.Println("Uncompleted tasks: ")
			todos.PrintUnCompleted()
		} else {
			todos.PrintAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	completed = listCmd.Flags().BoolP("completed", "c", false, "List all completed task")
	uncompleted = listCmd.Flags().BoolP("uncompleted", "u", false, "List all uncompleted completed task")
}
