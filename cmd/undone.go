package cmd

import (
	"github.com/spf13/cobra"
)

var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "unmark a todo as done",
	Long:  "unmark a todo as done by it's index, ex todo undone <id>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, idxStr := range args {
			todos.Undone(idxStr)
		}
		todos.PrintAll()
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
