package cmd

import (
	"github.com/spf13/cobra"
)

var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "unmark todos as done",
	Long:  "unmark todos as done by their index (allow multiple), ex: todo undone <id1> <id2>",
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
